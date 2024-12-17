package note

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Only allow a-z, A-Z, 0-9, '-', and '_' for valid note names
var filenameMatcher = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

// Note defines a struct to contain and wrap a text file that a user
// can edit. This note is stored as a markdown file with associated
// metadata. The note's content is assumed to be markdown and will be
// rendered as so
type Note struct {
	Metadata        // Note Metadata
	Content  string `json:"-"` // Note content (assumed to be markdown format)
}

// Generate and return markdown representation of the note. The note
// begins with yaml metadata and then contains markdown content after two
// new lines
func (a *Note) AsMarkdown() string {
	output := ""

	// Create yaml metadata at the top of the output
	output += "---\n"
	output += "author: " + a.Author + "\n"
	output += "createdAt: " + a.CreatedAt.Format(time.RFC3339) + "\n"
	output += "updatedAt: " + a.UpdatedAt.Format(time.RFC3339) + "\n"
	output += "---\n\n"

	// Add markdown content to the rest of the output
	output += a.Content

	return output
}

// Generate and return an HTML representation of the note. The note's
// metadata is rendered into a template, and then then note's markdown
// content is rendered as HTML using an external tool
func (a *Note) AsHTML() string {
	// Not implemented
	return ""
}

// Create a new note from a provided configuration and filename
func NewNote(config *Config, filename string) (*Note, error) {
	// Make sure the filename is valid
	if !filenameMatcher.MatchString(filename) {
		return nil, fmt.Errorf("invalid name '%s'", filename)
	}

	// Generate note title
	titleComponents := strings.Split(filename, "-")
	title := ""
	for _, t := range titleComponents {
		title += cases.Title(language.English).String(t) + " "
	}
	title = strings.Trim(title, " ")

	return &Note{
		Metadata: Metadata{
			Filename:  filename,
			Author:    config.DefaultAuthor,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Content: "# " + title + "\n\n",
	}, nil
}
