// 'publish' command saves a markdown file of the note to the provided directory
package main

import (
	"os"
	"path/filepath"

	"github.com/ethanbaker/note/pkg/note"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish [title] [directory|.]",
	Short: "Save a note to a directory",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		// Validate title input
		title := args[0]
		if title == "" {
			cmd.PrintErr("title cannot be empty")
			return
		}

		directory := "."
		if len(args) == 2 {
			// Validate directory input if provided
			directory = args[1]
			if directory == "" {
				cmd.PrintErr("directory cannot be empty")
				return
			}
		}

		// Get the note manager
		manager, err := note.GetManager()
		errHandler(cmd, err)

		// If note is nil, no note with the given title exists
		note := manager.GetNote(title)
		if note == nil {
			cmd.PrintErrf(`note "%s" does not exist\n`, title)
			return
		}

		// Save the note to the directory
		if err := os.WriteFile(filepath.Join(directory, note.Filename+".md"), []byte(note.AsMarkdown()), 0644); err != nil {
			errHandler(cmd, err)
		}

		// Print success message
		if directory == "." {
			cmd.Print("note saved to current directory\n")
		} else {
			cmd.Printf("note saved to %s\n", directory)
		}
	},
}
