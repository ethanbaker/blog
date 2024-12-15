package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

// Manager defines a struct to manage CRUD operations on articles
// using the config
type Manager struct {
	Articles []*Article `json:"articles"` // List of articles managed
	Config   *Config    `json:"-"`        // Config to manage articles
}

// Return status of if the manager contains the filename. If the manager contains the filename
// this method will return true and the index of the filename. If the manager does not contain
// the filename, this method will return false and -1
func (m *Manager) contains(filename string) (bool, int) {
	filename = strings.ToLower(filename)

	for i, article := range m.Articles {
		if article.Filename == filename {
			return true, i
		}
	}

	return false, -1
}

// Create a new article with the provided filename, save it to storage, and add it to the manager
func (m *Manager) CreateArticle(filename string) error {
	log.Printf("[INFO]: creating new article with filename '%s'", filename)

	filename = strings.ToLower(filename)

	// Check if a duplicate filename exists
	if ok, _ := m.contains(filename); ok {
		log.Printf("[ERR]: duplicate article name '%s'", filename)
		return fmt.Errorf("duplicate article name '%s'", filename)
	}

	// Create a new article
	article, err := NewArticle(m.Config, filename)
	if err != nil {
		log.Printf("[ERR]: failed to create new article (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully created new article with filename '%s'", filename)
	log.Printf("[INFO]: saving article to file")

	// Save the article's content to the filesystem
	filepath := path.Join(m.Config.Directory, filename+".md")
	log.Printf("[INFO]: saving article to file '%s'", filepath)

	if err = os.WriteFile(filepath, []byte(article.Content), 0666); err != nil {
		log.Printf("[ERR]: failed to save article to file (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully saved article to file")

	// Add the article to the manager
	m.Articles = append(m.Articles, article)

	log.Printf("[INFO]: saving manager")

	// Save the manager to storage
	if err := m.Save(); err != nil {
		log.Printf("[ERR]: failed to save manager (err: %v)", err)
		return err
	}
	return nil
}

// Delete an article with the provided filename, remove it from storage, and remove it from the manager
func (m *Manager) DeleteArticle(filename string) error {
	log.Printf("[INFO]: deleting article with filename '%s'", filename)

	filename = strings.ToLower(filename)

	// Find the article in the manager
	index, ok := -1, false
	if ok, index = m.contains(filename); !ok {
		log.Printf("[ERR]: article with name '%s' not found", filename)
		return fmt.Errorf("article with name '%s' not found", filename)
	}

	log.Printf("[INFO]: successfully found article with filename '%s', removing", filename)

	// Remove the article from the manager
	m.Articles = append(m.Articles[:index], m.Articles[index+1:]...)

	log.Printf("[INFO]: successfully removed article with filename '%s', deleting associated file", filename)

	// Remove the article from storage
	filepath := path.Join(m.Config.Directory, filename+".md")
	log.Printf("[INFO]: removing article at file '%s'", filepath)

	if err := os.Remove(filepath); err != nil {
		log.Printf("[ERR]: failed to remove article file (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully removed article file")
	log.Printf("[INFO]: saving manager")

	// Save the manager to storage
	if err := m.Save(); err != nil {
		log.Printf("[ERR]: failed to save manager (err: %v)", err)
		return err
	}

	return nil
}

// Open an article using the provided text editor
func (m *Manager) OpenArticle(filename string) error {
	log.Printf("[INFO]: opening article with filename '%s'", filename)

	filename = strings.ToLower(filename)

	// Make sure the filename exists in the manager
	index, ok := -1, false
	if ok, index = m.contains(filename); !ok {
		log.Printf("[ERR]: filename '%s' does not exist", filename)
		return fmt.Errorf("article with name '%s' not found", filename)
	}

	log.Printf("[INFO]: found article with filename '%s'", filename)
	log.Printf("[INFO]: opening article in editor")

	// Get article details
	article := m.Articles[index]
	filepath := path.Join(m.Config.Directory, filename+".md")

	log.Printf("[INFO]: getting article details at file '%s'", filepath)

	// Open the file in the editor
	if err := exec.Command(m.Config.Editor, filepath).Run(); err != nil {
		log.Printf("[ERR]: failed to open article in editor (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: article edited, continuing")
	log.Printf("[INFO]: updating article metadata")

	// Save the article with the manager
	article.UpdatedAt = time.Now()

	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("[ERR]: failed to read article file (err: %v)", err)
		return err
	}
	article.Content = string(content)

	log.Printf("[INFO]: saving manager")

	// Save the manager to storage
	if err := m.Save(); err != nil {
		log.Printf("[ERR]: failed to save manager (err: %v)", err)
		return err
	}

	return nil
}

// Return a list of all articles in the manager
func (m *Manager) GetArticles() []*Article {
	log.Printf("[INFO]: listing all articles")

	return m.Articles
}

// Return an article with the provided filename. If no matching article can
// be found, this method will return nil
func (m *Manager) GetArticle(filename string) *Article {
	log.Printf("[INFO]: getting article with filename '%s'", filename)

	for _, article := range m.Articles {
		if article.Filename == filename {
			log.Printf("[INFO]: found article with filename '%s'", filename)
			return article
		}
	}

	return nil
}

// Save all article-related metadata to storage
func (m *Manager) Save() error {
	log.Printf("[INFO]: saving manager information")

	// Save all article metadata
	file, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		log.Printf("[ERR]: failed to marshal manager struct (err: %v)", err)
		return err
	}

	// Write the JSON object to the default filepath
	if err = os.WriteFile(managerPath, file, 0666); err != nil {
		log.Printf("[ERR]: failed to save manager struct to file (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully saved manager struct to file")
	log.Printf("[INFO]: saving articles to files")

	// Save all article content to a file represented by their filename
	for _, article := range m.Articles {
		log.Printf("[INFO]: saving article '%s' to file", article.Filename)

		filepath := path.Join(m.Config.Directory, article.Filename+".md")
		log.Printf("[INFO]: saving article to file '%s'", filepath)

		if err := os.WriteFile(filepath, []byte(article.Content), 0666); err != nil {
			log.Printf("[ERR]: failed to save article '%s' to file (err: %v)", article.Filename, err)
			return err
		}
		log.Printf("[INFO]: successfully saved article '%s' to file", article.Filename)
	}

	log.Printf("[INFO]: successfully saved articles to files")
	log.Printf("[INFO]: saving config file")

	// Save config file
	if err := m.Config.Save(); err != nil {
		log.Printf("[ERR]: failed to save config file (err: %v)", err)
		return err
	}

	return nil
}

// Load related article metadata from storage
func (m *Manager) Load() error {
	log.Printf("[INFO]: reading manager information")

	// Read in the provided filename
	file, err := os.ReadFile(managerPath)
	if err != nil {
		log.Printf("[ERR]: failed to read manager file (err: %v)", err)
		return err
	}

	// Unmarshal the file into the Manager struct
	if err = json.Unmarshal(file, m); err != nil {
		log.Printf("[ERR]: failed to parse manager file (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully read into manager struct")
	log.Printf("[INFO]: reading articles from files")

	// Read all article content from files
	for _, article := range m.Articles {
		log.Printf("[INFO]: reading article '%s' from file", article.Filename)

		filepath := path.Join(m.Config.Directory, article.Filename+".md")
		log.Printf("[INFO]: reading article from file '%s'", filepath)

		content, err := os.ReadFile(filepath)
		if err != nil {
			log.Printf("[ERR]: failed to read article '%s' from file (err: %v)", article.Filename, err)
			return err
		}
		article.Content = string(content)

		log.Printf("[INFO]: successfully read article '%s' from file", article.Filename)
	}

	log.Printf("[INFO]: successfully read articles from files")
	log.Printf("[INFO]: reading config file")

	// Read config file
	if err := m.Config.Load(); err != nil {
		log.Printf("[ERR]: failed to read config file (err: %v)", err)
		return err
	}

	return nil
}

// GetManager returns an active instance of the manager with the stored config
func GetManager() (*Manager, error) {
	manager := &Manager{
		Config: NewConfig(),
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		// Create the config.json file if it doesn't already exist
		log.Printf("[INFO]: config.json file does not exist, creating it")

		if err := manager.Config.Save(); err != nil {
			log.Printf("[ERR]: failed to save config (err: %v)", err)
			return nil, err
		}
	} else {
		// The config file already exists, so load it
		log.Printf("[INFO]: config.json file exists, loading it")

		if err := manager.Config.Load(); err != nil {
			log.Printf("[ERR]: failed to load config (err: %v)", err)
			return nil, err
		}
	}

	if _, err := os.Stat(managerPath); errors.Is(err, os.ErrNotExist) {
		// Create the manager.json file if it doesn't already exist
		log.Printf("[INFO]: manager.json file does not exist, creating it")

		if err := manager.Save(); err != nil {
			log.Printf("[ERR]: failed to create manager.json file")
			return nil, err
		}
	} else {
		// The manager file already exists, so load it
		log.Printf("[INFO]: manager.json file exists, loading it")

		if err := manager.Load(); err != nil {
			log.Printf("[ERR]: failed to load manager (err: %v)", err)
			return nil, err
		}
	}

	log.Printf("[INFO]: checking existance of article directory")

	// Create the directory where articles are stored if it doesn't already exist
	if _, err := os.Stat(manager.Config.Directory); errors.Is(err, os.ErrNotExist) {
		log.Printf("[INFO]: article directory file does not exist, creating it")

		if err = os.Mkdir(manager.Config.Directory, 0755); err != nil {
			log.Printf("[ERR]: failed to create article directory file")
			return nil, err
		}
	}

	return manager, nil
}
