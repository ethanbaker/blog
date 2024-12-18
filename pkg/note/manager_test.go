package note_test

import (
	"errors"
	"os"
	"path"
	"testing"

	"github.com/ethanbaker/note/pkg/note"
	"github.com/stretchr/testify/require"
)

// Setup before each tests by initializing a config struct
func managerTestSetup() (*note.Manager, error) {
	// Get working directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Clear the testing directory
	testingDir := path.Join(wd, "testing/dirty")
	if err := os.RemoveAll(testingDir); err != nil {
		return nil, err
	}
	if err := os.Mkdir(testingDir, 0755); err != nil {
		return nil, err
	}

	// Modify default paths to test files
	note.ModifyDefaultDirectoryPath(path.Join(wd, "testing/dirty/entries/"))
	note.ModifyConfigPath(path.Join(wd, "testing/dirty/config.json"))
	note.ModifyManagerPath(path.Join(wd, "testing/dirty/manager.json"))

	// Get the default manager
	manager, err := note.GetManager()
	if err != nil {
		return nil, err
	}

	// Add testing config options to the manager
	manager.Config.Editor = "cat" // Dummy editor to spit out file contents instead of actually edit
	manager.Config.DefaultAuthor = "Ethan"

	return manager, nil
}

// Test creating a new note
func TestCreateNote(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new note
	require.Nil(manager.CreateNote("note-1"))
}

// Test creating a duplicate note
func TestCreateNoteDuplicate(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new note
	require.Nil(manager.CreateNote("note-1"))

	// Create a duplicate note
	err = manager.CreateNote("note-1")
	require.NotNil(err)
	require.Equal("duplicate note name 'note-1'", err.Error())
}

// Test creating an invalid note
func TestCreateNoteInvalid(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create an invalid note
	err = manager.CreateNote("$note")
	require.NotNil(err)
	require.Equal("invalid name '$note'", err.Error())
}

// Test getting a single note
func TestGetNote(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new note
	require.Nil(manager.CreateNote("note-1"))

	// Get the note
	note := manager.GetNote("note-1")
	require.NotNil(note)

	require.Equal("# Note 1\n\n", note.Content)
	require.Equal("note-1", note.Filename)
	require.Equal("Ethan", note.Author)
	require.NotNil(note.CreatedAt)
	require.NotNil(note.UpdatedAt)

	// Check the content of the associated note file
	content, err := os.ReadFile("./testing/dirty/entries/note-1.md")
	require.Nil(err)
	require.Equal("# Note 1\n\n", string(content))
}

// Test getting multiple notes
func TestGetNotes(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new note
	require.Nil(manager.CreateNote("note-1"))
	require.Nil(manager.CreateNote("note-2"))

	// Get all notes
	notes := manager.GetNotes()
	require.Len(notes, 2)

	// Check the first note
	require.Equal("# Note 1\n\n", notes[0].Content)
	require.Equal("note-1", notes[0].Filename)
	require.Equal("Ethan", notes[0].Author)
	require.NotNil(notes[0].CreatedAt)
	require.NotNil(notes[0].UpdatedAt)

	// Check the second note
	require.Equal("# Note 2\n\n", notes[1].Content)
	require.Equal("note-2", notes[1].Filename)
	require.Equal("Ethan", notes[1].Author)
	require.NotNil(notes[1].CreatedAt)
	require.NotNil(notes[1].UpdatedAt)

	// Check the content of the associated note files
	content, err := os.ReadFile("./testing/dirty/entries/note-1.md")
	require.Nil(err)
	require.Equal("# Note 1\n\n", string(content))

	content, err = os.ReadFile("./testing/dirty/entries/note-2.md")
	require.Nil(err)
	require.Equal("# Note 2\n\n", string(content))
}

// Test deleting an note
func TestDeleteNote(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new note
	require.Nil(manager.CreateNote("note-1"))

	// Get the note to verify it's been saved
	note := manager.GetNote("note-1")
	require.NotNil(note)

	require.Equal("# Note 1\n\n", note.Content)
	require.Equal("note-1", note.Filename)
	require.Equal("Ethan", note.Author)
	require.NotNil(note.CreatedAt)
	require.NotNil(note.UpdatedAt)

	// Check the content of the associated note file
	content, err := os.ReadFile("./testing/dirty/entries/note-1.md")
	require.Nil(err)
	require.Equal("# Note 1\n\n", string(content))

	// Delete the note
	require.Nil(manager.DeleteNote("note-1"))

	// Verify the note is gone
	require.Nil(manager.GetNote("note-1"))

	// Verify the note file is gone
	_, err = os.Stat("./testing/dirty/entries/note-1.md")
	require.True(errors.Is(err, os.ErrNotExist))
}

// Test deleting an note that doesn't exist
func TestDeleteNoteNotFound(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Delete the note that doesn't exist
	err = manager.DeleteNote("note-1")
	require.NotNil(err)
	require.Equal("note with name 'note-1' not found", err.Error())
}

// Test opening an note
func TestOpenNote(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new note
	require.Nil(manager.CreateNote("note-1"))

	// Open the note
	require.Nil(manager.OpenNote("note-1"))
}

// Test opening an note that doesn't exist
func TestOpenNoteNotFound(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Delete the note that doesn't exist
	err = manager.OpenNote("note-1")
	require.NotNil(err)
	require.Equal("note with name 'note-1' not found", err.Error())
}

// Test loading a manager that already exists
func TestLoad(t *testing.T) {
	require := require.New(t)

	// Get working directory
	wd, err := os.Getwd()
	require.Nil(err)

	// Modify default paths to test files
	note.ModifyDefaultDirectoryPath(path.Join(wd, "testing/pristine/notes/entries/"))
	note.ModifyConfigPath(path.Join(wd, "testing/pristine/config.json"))
	note.ModifyManagerPath(path.Join(wd, "testing/pristine/manager.json"))

	// Get the default manager
	manager, err := note.GetManager()
	require.Nil(err)

	// Check manager config
	require.Equal("testing/pristine/entries/", manager.Config.Directory)
	require.Equal("vi", manager.Config.Editor)
	require.Equal("John", manager.Config.DefaultAuthor)

	// Check note list
	require.Equal(2, len(manager.Notes))

	note1 := manager.Notes[0]
	require.Equal("note-1", note1.Filename)
	require.Equal("John", note1.Author)
	require.NotNil(note1.CreatedAt)
	require.NotNil(note1.UpdatedAt)

	note2 := manager.Notes[1]
	require.Equal("note-2", note2.Filename)
	require.Equal("Mary", note2.Author)
	require.NotNil(note2.CreatedAt)
	require.NotNil(note2.UpdatedAt)
}
