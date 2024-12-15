package cli

import "time"

// Metadata contains generic metadata for an article
type Metadata struct {
	Filename  string    `json:"filename"`  // Filename of the article (used to associate where the article is stored)
	Author    string    `json:"author"`    // The author of the article
	CreatedAt time.Time `json:"createdAt"` // Time the article was last created
	UpdatedAt time.Time `json:"updatedAt"` // The the article was last updated
}
