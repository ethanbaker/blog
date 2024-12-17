package note

import "time"

// Metadata contains generic metadata for an note
type Metadata struct {
	Filename  string    `json:"filename"`  // Filename of the note (used to associate where the note is stored)
	Author    string    `json:"author"`    // The author of the note
	CreatedAt time.Time `json:"createdAt"` // Time the note was last created
	UpdatedAt time.Time `json:"updatedAt"` // The the note was last updated
}
