// 'remove' command removes existing notes
package main

import (
	"github.com/ethanbaker/note/pkg/note"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove [title]",
	Short:   "Removes an note",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"rm"},
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

		// Remove the note
		err = manager.DeleteNote(title)
		errHandler(cmd, err)

		// Print success message
		cmd.Println("note removed successfully")
	},
}
