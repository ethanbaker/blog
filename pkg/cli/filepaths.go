package cli

// Path to the directory where articles are stored
var defaultDirectoryPath = "~/.local/share/notes/articles/"

// Path to the configuration file for the system
var configPath = "~/.config/note.json"

// Path to the manager file where article metadata is stored
var managerPath = "~/.local/share/notes/manager.json"

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
