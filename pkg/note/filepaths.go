package note

import (
	"os"
	"path"
)

// Path to the directory where notes are stored
var defaultDirectoryPath string

// Path to the configuration file for the system
var configPath string

// Path to the manager file where note metadata is stored
var managerPath string

// Get the user's home directory to concatenate with default paths
func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	defaultDirectoryPath = path.Join(home, ".local/share/notes/entries/")
	configPath = path.Join(home, ".config/note.json")
	managerPath = path.Join(home, ".local/share/notes/manager.json")
}

// Change the default directory path. This is an experimental function and
// should be avoided unless absolutely necessary
func ModifyDefaultDirectoryPath(path string) {
	defaultDirectoryPath = path
}

// Change the config path. This is an experimental function and should
// be avoided unless absolutely necessary
func ModifyConfigPath(path string) {
	configPath = path
}

// Modify the manager path. This is an experimental function and should
// be avoided unless absolutely necessary
func ModifyManagerPath(path string) {
	managerPath = path
}
