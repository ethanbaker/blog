package note_test

import (
	"strings"
	"testing"
	"time"

	"github.com/ethanbaker/note/pkg/note"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Setup before each tests by initializing a config struct
func noteTestSetup() *note.Config {
	return &note.Config{
		Directory:     "./testing/dirty/entries/",
		Editor:        "vi",
		DefaultAuthor: "Ethan",
	}
}

// TestNewNote tests the creation of a new note
func TestNewNote(t *testing.T) {
	// Setup testing objects
	require := require.New(t)
	config := noteTestSetup()

	// Create a new note
	note, err := note.NewNote(config, "test-note")
	require.Nil(err)

	// Test note metadata
	require.Equal("Ethan", note.Author)
	require.Equal("test-note", note.Filename)
	require.NotEmpty(note.CreatedAt)
	require.NotEmpty(note.UpdatedAt)

	// Test note content
	require.Equal("# Test Note\n\n", note.Content)
}

// TestNoteAsMarkdown tests the generation of an note represented in markdown
func TestNoteAsMarkdown(t *testing.T) {
	// Setup testing objects
	require := require.New(t)
	assert := assert.New(t)
	config := noteTestSetup()

	// Create a new note
	note, err := note.NewNote(config, "test-note")
	require.Nil(err)

	// Test markdown representation
	markdown := note.AsMarkdown()
	lines := strings.Split(markdown, "\n")

	require.Equal(9, len(lines))

	// Check line equality
	assert.Equal("---", lines[0])
	assert.Equal("author: Ethan", lines[1])
	assert.Equal("createdAt: "+note.CreatedAt.Format(time.RFC3339), lines[2])
	assert.Equal("updatedAt: "+note.UpdatedAt.Format(time.RFC3339), lines[3])
	assert.Equal("---", lines[4])
	assert.Equal("", lines[5])
	assert.Equal("# Test Note", lines[6])
	assert.Equal("", lines[7])
	assert.Equal("", lines[8])
}
