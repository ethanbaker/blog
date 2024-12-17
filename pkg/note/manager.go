package note

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

// Manager defines a struct to manage CRUD operations on notes
// using the config
type Manager struct {
	Notes  []*Note `json:"notes"` // List of notes managed
	Config *Config `json:"-"`     // Config to manage notes
}

// Return status of if the manager contains the filename. If the manager contains the filename
// this method will return true and the index of the filename. If the manager does not contain
// the filename, this method will return false and -1
func (m *Manager) contains(filename string) (bool, int) {
	filename = strings.ToLower(filename)

	for i, note := range m.Notes {
		if note.Filename == filename {
			return true, i
		}
	}

	return false, -1
}

// Create a new note with the provided filename, save it to storage, and add it to the manager
func (m *Manager) CreateNote(filename string) error {
	log.Printf("[INFO]: creating new note with filename '%s'", filename)

	filename = strings.ToLower(filename)

	// Check if a duplicate filename exists
	if ok, _ := m.contains(filename); ok {
		log.Printf("[ERR]: duplicate note name '%s'", filename)
		return fmt.Errorf("duplicate note name '%s'", filename)
	}

	// Create a new note
	note, err := NewNote(m.Config, filename)
	if err != nil {
		log.Printf("[ERR]: failed to create new note (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully created new note with filename '%s'", filename)
	log.Printf("[INFO]: saving note to file")

	// Save the note's content to the filesystem
	filepath := path.Join(m.Config.Directory, filename+".md")
	log.Printf("[INFO]: saving note to file '%s'", filepath)

	if err = os.WriteFile(filepath, []byte(note.Content), 0600); err != nil {
		log.Printf("[ERR]: failed to save note to file (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully saved note to file")

	// Add the note to the manager
	m.Notes = append(m.Notes, note)

	log.Printf("[INFO]: saving manager")

	// Save the manager to storage
	if err := m.Save(); err != nil {
		log.Printf("[ERR]: failed to save manager (err: %v)", err)
		return err
	}
	return nil
}

// Delete an note with the provided filename, remove it from storage, and remove it from the manager
func (m *Manager) DeleteNote(filename string) error {
	log.Printf("[INFO]: deleting note with filename '%s'", filename)

	filename = strings.ToLower(filename)

	// Find the note in the manager
	index, ok := -1, false
	if ok, index = m.contains(filename); !ok {
		log.Printf("[ERR]: note with name '%s' not found", filename)
		return fmt.Errorf("note with name '%s' not found", filename)
	}

	log.Printf("[INFO]: successfully found note with filename '%s', removing", filename)

	// Remove the note from the manager
	m.Notes = append(m.Notes[:index], m.Notes[index+1:]...)

	log.Printf("[INFO]: successfully removed note with filename '%s', deleting associated file", filename)

	// Remove the note from storage
	filepath := path.Join(m.Config.Directory, filename+".md")
	log.Printf("[INFO]: removing note at file '%s'", filepath)

	if err := os.Remove(filepath); err != nil {
		log.Printf("[ERR]: failed to remove note file (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully removed note file")
	log.Printf("[INFO]: saving manager")

	// Save the manager to storage
	if err := m.Save(); err != nil {
		log.Printf("[ERR]: failed to save manager (err: %v)", err)
		return err
	}

	return nil
}

// Open an note using the provided text editor
func (m *Manager) OpenNote(filename string) error {
	log.Printf("[INFO]: opening note with filename '%s'", filename)

	filename = strings.ToLower(filename)

	// Make sure the filename exists in the manager
	index, ok := -1, false
	if ok, index = m.contains(filename); !ok {
		log.Printf("[ERR]: filename '%s' does not exist", filename)
		return fmt.Errorf("note with name '%s' not found", filename)
	}

	log.Printf("[INFO]: found note with filename '%s'", filename)
	log.Printf("[INFO]: opening note in editor")

	// Get note details
	note := m.Notes[index]
	filepath := path.Join(m.Config.Directory, filename+".md")

	log.Printf("[INFO]: getting note details at file '%s'", filepath)
	log.Printf("[INFO]: command = '%s %s'", m.Config.Editor, filepath)

	// Create editor command
	cmd := exec.Command(m.Config.Editor, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Open the note in the editor
	if err := cmd.Run(); err != nil {
		log.Printf("[ERR]: failed to open note in editor (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: note edited, continuing")
	log.Printf("[INFO]: updating note metadata")

	// Save the note with the manager
	note.UpdatedAt = time.Now()

	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("[ERR]: failed to read note file (err: %v)", err)
		return err
	}
	note.Content = string(content)

	log.Printf("[INFO]: saving manager")

	// Save the manager to storage
	if err := m.Save(); err != nil {
		log.Printf("[ERR]: failed to save manager (err: %v)", err)
		return err
	}

	return nil
}

// Return a list of all notes in the manager
func (m *Manager) GetNotes() []*Note {
	log.Printf("[INFO]: listing all notes")

	return m.Notes
}

// Return an note with the provided filename. If no matching note can
// be found, this method will return nil
func (m *Manager) GetNote(filename string) *Note {
	log.Printf("[INFO]: getting note with filename '%s'", filename)

	for _, note := range m.Notes {
		if note.Filename == filename {
			log.Printf("[INFO]: found note with filename '%s'", filename)
			return note
		}
	}

	return nil
}

// Save all note-related metadata to storage
func (m *Manager) Save() error {
	log.Printf("[INFO]: saving manager information")

	// Save all note metadata
	file, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		log.Printf("[ERR]: failed to marshal manager struct (err: %v)", err)
		return err
	}

	// Ensure the directory exists
	err = os.MkdirAll(path.Dir(managerPath), 0755)
	if err != nil {
		log.Printf("[ERR]: failed to create config directory (err: %v)", err)
		return err
	}

	// Write the JSON object to the default filepath
	if err = os.WriteFile(managerPath, file, 0600); err != nil {
		log.Printf("[ERR]: failed to save manager struct to file (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully saved manager struct to file")
	log.Printf("[INFO]: saving notes to files")

	// Save all note content to a file represented by their filename
	for _, note := range m.Notes {
		log.Printf("[INFO]: saving note '%s' to file", note.Filename)

		filepath := path.Join(m.Config.Directory, note.Filename+".md")
		log.Printf("[INFO]: saving note to file '%s'", filepath)

		if err := os.WriteFile(filepath, []byte(note.Content), 0600); err != nil {
			log.Printf("[ERR]: failed to save note '%s' to file (err: %v)", note.Filename, err)
			return err
		}
		log.Printf("[INFO]: successfully saved note '%s' to file", note.Filename)
	}

	log.Printf("[INFO]: successfully saved notes to files")
	log.Printf("[INFO]: saving config file")

	// Save config file
	if err := m.Config.Save(); err != nil {
		log.Printf("[ERR]: failed to save config file (err: %v)", err)
		return err
	}

	return nil
}

// Load related note metadata from storage
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
	log.Printf("[INFO]: reading notes from files")

	// Read all note content from files
	for _, note := range m.Notes {
		log.Printf("[INFO]: reading note '%s' from file", note.Filename)

		filepath := path.Join(m.Config.Directory, note.Filename+".md")
		log.Printf("[INFO]: reading note from file '%s'", filepath)

		content, err := os.ReadFile(filepath)
		if err != nil {
			log.Printf("[ERR]: failed to read note '%s' from file (err: %v)", note.Filename, err)
			return err
		}
		note.Content = string(content)

		log.Printf("[INFO]: successfully read note '%s' from file", note.Filename)
	}

	log.Printf("[INFO]: successfully read notes from files")
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

	log.Printf("[INFO]: checking existance of note directory")

	// Create the directory where notes are stored if it doesn't already exist
	if _, err := os.Stat(manager.Config.Directory); errors.Is(err, os.ErrNotExist) {
		log.Printf("[INFO]: note directory file does not exist, creating it")

		if err = os.MkdirAll(manager.Config.Directory, 0755); err != nil {
			log.Printf("[ERR]: failed to create note directory file")
			return nil, err
		}
	}

	return manager, nil
}
