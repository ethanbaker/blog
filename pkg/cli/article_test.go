package cli_test

import (
	"strings"
	"testing"

	"github.com/ethanbaker/note/pkg/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Setup before each tests by initializing a config struct
func articleTestSetup() *cli.Config {
	return &cli.Config{
		Directory:     "./testing/dirty/articles/",
		Editor:        "vi",
		DefaultAuthor: "Ethan",
	}
}

// TestNewArticle tests the creation of a new article
func TestNewArticle(t *testing.T) {
	// Setup testing objects
	require := require.New(t)
	config := articleTestSetup()

	// Create a new article
	article, err := cli.NewArticle(config, "test-article")
	require.Nil(err)

	// Test article metadata
	require.Equal("Ethan", article.Author)
	require.Equal("test-article", article.Filename)
	require.NotEmpty(article.CreatedAt)
	require.NotEmpty(article.UpdatedAt)

	// Test article content
	require.Equal("# Test Article\n\n", article.Content)
}

// TestArticleAsMarkdown tests the generation of an article represented in markdown
func TestArticleAsMarkdown(t *testing.T) {
	// Setup testing objects
	require := require.New(t)
	assert := assert.New(t)
	config := articleTestSetup()

	// Create a new article
	article, err := cli.NewArticle(config, "test-article")
	require.Nil(err)

	// Test markdown representation
	markdown := article.AsMarkdown()
	lines := strings.Split(markdown, "\n")

	require.Equal(9, len(lines))

	// Check line equality
	assert.Equal("---", lines[0])
	assert.Equal("author: Ethan", lines[1])
	assert.Equal("createdAt: "+article.CreatedAt.Format("2006-01-02T15:04:05-07:00"), lines[2])
	assert.Equal("updatedAt: "+article.UpdatedAt.Format("2006-01-02T15:04:05-07:00"), lines[3])
	assert.Equal("---", lines[4])
	assert.Equal("", lines[5])
	assert.Equal("# Test Article", lines[6])
	assert.Equal("", lines[7])
	assert.Equal("", lines[8])
}
