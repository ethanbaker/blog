package note_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test opening the config file
func TestOpenConfig(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new note
	require.Nil(manager.OpenConfig())
}
