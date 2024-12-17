// 'new' command makes a new note
package main

import (
	"github.com/ethanbaker/note/pkg/note"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [title]",
	Short: "Create a new note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Validate title input
		title := args[0]
		if title == "" {
			cmd.PrintErr("Title cannot be empty")
			return
		}

		// Get the note manager
		manager, err := note.GetManager()
		errHandler(cmd, err)

		// Create a new note
		err = manager.CreateNote(title)
		errHandler(cmd, err)

		// Open the newly created note
		err = manager.OpenNote(title)
		errHandler(cmd, err)

		// Print success message
		cmd.Println("note created successfully")
	},
}
