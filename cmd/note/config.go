// 'config' command to manage configuration settings
package main

import (
	"github.com/ethanbaker/note/pkg/note"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Open the configuration JSON file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the note manager
		manager, err := note.GetManager()
		errHandler(cmd, err)

		// Open the configuration file
		err = manager.OpenConfig()
		errHandler(cmd, err)

		// Print success message
		cmd.Println("configuration file edited successfully")
	},
}
