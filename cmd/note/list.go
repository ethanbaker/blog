// 'list' command lists existing notes
package main

import (
	"fmt"
	"text/tabwriter"

	"github.com/ethanbaker/note/pkg/note"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List existing notes",
	Aliases: []string{"ls"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the note manager
		manager, err := note.GetManager()
		errHandler(cmd, err)

		// Get all notes
		notes := manager.GetNotes()

		// If there are no notes, print a message and return
		if len(notes) == 0 {
			cmd.Println("No notes found")
			return
		}

		// Otherwise, print all notes as a table
		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 2, 4, ' ', 0)
		fmt.Fprintf(w, "FILENAME\tCREATED ON\tLAST UPDATED\n")

		for _, note := range notes {
			fmt.Fprintf(w, "%s\t%s\t%s\n", note.Filename, note.CreatedAt.Format("2006-01-02"), note.UpdatedAt.Format("2006-01-02"))
		}

		if err := w.Flush(); err != nil {
			errHandler(cmd, err)
		}
	},
}
