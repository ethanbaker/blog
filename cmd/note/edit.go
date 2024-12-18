// 'edit' command to open an existing note
package main

import (
	"github.com/ethanbaker/note/pkg/note"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [title]",
	Short: "Open and edit an existing note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Validate title input
		title := args[0]
		if title == "" {
			cmd.PrintErr("title cannot be empty")
			return
		}

		// Get the note manager
		manager, err := note.GetManager()
		errHandler(cmd, err)

		// Open the note
		err = manager.OpenNote(title)
		errHandler(cmd, err)

		// Print success message
		cmd.Printf("note \"%s\" saved successfully\n", title)
	},
}
