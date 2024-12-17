package note

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

// Config struct represents the tool's static configuration that is loaded
// from a JSON file
type Config struct {
	Directory     string `json:"directory"`      // Directory where notes are stored
	Editor        string `json:"editor"`         // Editor for opening notes, represented as a command
	DefaultAuthor string `json:"default_author"` // Default author for new notes
}

// Return a copy of the existing config
func (c *Config) Copy() *Config {
	return &Config{
		Directory:     c.Directory,
		Editor:        c.Editor,
		DefaultAuthor: c.DefaultAuthor,
	}
}

// Load the configuration from the default filepath and load it into an empty struct
func (c *Config) Load() error {
	log.Printf("[INFO]: loading config file")

	// Read in the provided filename
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("[ERR]: failed to read config file (err: %v)", err)
		return err
	}

	// Unmarshal the file into the Config struct
	if err = json.Unmarshal(file, c); err != nil {
		log.Printf("[ERR]: failed to parse config file (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully loaded config file")
	return nil
}

// Save the configuration to the specified default filepath
func (c *Config) Save() error {
	// Marshal the config struct into a JSON object
	file, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Printf("[ERR]: failed to marshal config file (err: %v)", err)
		return err
	}

	// Ensure the directory exists
	err = os.MkdirAll(path.Dir(configPath), 0755)
	if err != nil {
		log.Printf("[ERR]: failed to create config directory (err: %v)", err)
		return err
	}

	// Write the JSON object to the default filepath
	if err = os.WriteFile(configPath, file, 0600); err != nil {
		log.Printf("[ERR]: failed to save config file (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: successfully saved config file")
	return nil
}

// Create a new configuration with default values
func NewConfig() *Config {
	return &Config{
		Directory:     defaultDirectoryPath,
		Editor:        "vi",
		DefaultAuthor: "Anonymous",
	}
}

// Open the existing config file
func (m *Manager) OpenConfig() error {
	log.Printf("[INFO]: opening config file")

	// Make sure the config file exists in the manager
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Printf("[ERR]: config file does not exist")
		return fmt.Errorf("config file does not exist")
	}

	log.Printf("[INFO]: opening config file in editor")

	// Save existing config file state in case user inputs invalid data
	old := m.Config.Copy()

	// Create open command
	cmd := exec.Command(m.Config.Editor, configPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Open the file in the editor
	if err := cmd.Run(); err != nil {
		log.Printf("[ERR]: failed to open config file in editor (err: %v)", err)
		return err
	}

	log.Printf("[INFO]: config edited, continuing")
	log.Printf("[INFO]: validating config file")

	// Validating config file
	if err := m.Config.Load(); err != nil {
		log.Printf("[ERR]: failed to load updated config file (err: %v)", err)

		// On error, revert to the old config
		m.Config = old
		m.Config.Save()

		return err
	}

	return nil
}
