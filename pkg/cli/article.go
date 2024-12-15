package cli

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Only allow a-z, A-Z, 0-9, '-', and '_' for valid article names
var filenameMatcher = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

// Article defines a struct to contain and wrap a text file that a user
// can edit. This article is stored as a markdown file with associated
// metadata. The article's content is assumed to be markdown and will be
// rendered as so
type Article struct {
	Metadata        // Article Metadata
	Content  string `json:"-"` // Article content (assumed to be markdown format)
}

// Generate and return markdown representation of the article. The article
// begins with yaml metadata and then contains markdown content after two
// new lines
func (a *Article) AsMarkdown() string {
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

// Generate and return an HTML representation of the article. The article's
// metadata is rendered into a template, and then then article's markdown
// content is rendered as HTML using an external tool
func (a *Article) AsHTML() string {
	// Not implemented
	return ""
}

// Create a new article from a provided configuration and filename
func NewArticle(config *Config, filename string) (*Article, error) {
	// Make sure the filename is valid
	if !filenameMatcher.MatchString(filename) {
		return nil, fmt.Errorf("invalid name '%s'", filename)
	}

	// Generate article title
	titleComponents := strings.Split(filename, "-")
	title := ""
	for _, t := range titleComponents {
		title += cases.Title(language.English).String(t) + " "
	}
	title = strings.Trim(title, " ")

	return &Article{
		Metadata: Metadata{
			Filename:  filename,
			Author:    config.DefaultAuthor,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Content: "# " + title + "\n\n",
	}, nil
}
