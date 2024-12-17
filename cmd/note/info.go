// 'info' command returns information about an note
package main

import (
	"fmt"
	"text/tabwriter"

	"github.com/ethanbaker/note/pkg/note"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info [title]",
	Short: "Show metadata about a note",
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

		// Get the note
		note := manager.GetNote(title)
		errHandler(cmd, err)

		// If note is nil, no note with the given title exists
		if note == nil {
			cmd.PrintErrf(`note "%s" does not exist\n`, title)
			return
		}

		// Print note metadata
		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 2, 2, ' ', 0)

		fmt.Fprintf(w, "FILENAME\t%s\n", note.Filename)
		fmt.Fprintf(w, "CREATED ON\t%s\n", note.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Fprintf(w, "LAST UPDATED\t%s\n", note.UpdatedAt.Format("2006-01-02 15:04:05"))

		if err := w.Flush(); err != nil {
			errHandler(cmd, err)
		}
	},
}
