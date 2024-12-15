package cli

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct represents the tool's static configuration that is loaded
// from a JSON file
type Config struct {
	Directory     string `json:"directory"`     // Directory where articles are stored
	Editor        string `json:"editor"`        // Editor for opening articles, represented as a command
	DefaultAuthor string `json:"defaultAuthor"` // Default author for new articles
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

	// Write the JSON object to the default filepath
	if err = os.WriteFile(configPath, file, 0666); err != nil {
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
