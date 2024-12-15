package cli_test

import (
	"errors"
	"os"
	"path"
	"testing"

	"github.com/ethanbaker/note/pkg/cli"
	"github.com/stretchr/testify/require"
)

// Setup before each tests by initializing a config struct
func managerTestSetup() (*cli.Manager, error) {
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
	cli.ModifyDefaultDirectoryPath(path.Join(wd, "testing/dirty/articles/"))
	cli.ModifyConfigPath(path.Join(wd, "testing/dirty/config.json"))
	cli.ModifyManagerPath(path.Join(wd, "testing/dirty/manager.json"))

	// Get the default manager
	manager, err := cli.GetManager()
	if err != nil {
		return nil, err
	}

	// Add testing config options to the manager
	manager.Config.Editor = "cat" // Dummy editor to spit out file contents instead of actually edit
	manager.Config.DefaultAuthor = "Ethan"

	return manager, nil
}

// Test creating a new article
func TestCreateArticle(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new article
	require.Nil(manager.CreateArticle("article-1"))
}

// Test creating a duplicate article
func TestCreateArticleDuplicate(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new article
	require.Nil(manager.CreateArticle("article-1"))

	// Create a duplicate article
	err = manager.CreateArticle("article-1")
	require.NotNil(err)
	require.Equal("duplicate article name 'article-1'", err.Error())
}

// Test creating an invalid article
func TestCreateArticleInvalid(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create an invalid article
	err = manager.CreateArticle("$article")
	require.NotNil(err)
	require.Equal("invalid name '$article'", err.Error())
}

// Test getting a single article
func TestGetArticle(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new article
	require.Nil(manager.CreateArticle("article-1"))

	// Get the article
	article := manager.GetArticle("article-1")
	require.NotNil(article)

	require.Equal("# Article 1\n\n", article.Content)
	require.Equal("article-1", article.Filename)
	require.Equal("Ethan", article.Author)
	require.NotNil(article.CreatedAt)
	require.NotNil(article.UpdatedAt)

	// Check the content of the associated article file
	content, err := os.ReadFile("./testing/dirty/articles/article-1.md")
	require.Nil(err)
	require.Equal("# Article 1\n\n", string(content))
}

// Test getting multiple articles
func TestGetArticles(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new article
	require.Nil(manager.CreateArticle("article-1"))
	require.Nil(manager.CreateArticle("article-2"))

	// Get all articles
	articles := manager.GetArticles()
	require.Len(articles, 2)

	// Check the first article
	require.Equal("# Article 1\n\n", articles[0].Content)
	require.Equal("article-1", articles[0].Filename)
	require.Equal("Ethan", articles[0].Author)
	require.NotNil(articles[0].CreatedAt)
	require.NotNil(articles[0].UpdatedAt)

	// Check the second article
	require.Equal("# Article 2\n\n", articles[1].Content)
	require.Equal("article-2", articles[1].Filename)
	require.Equal("Ethan", articles[1].Author)
	require.NotNil(articles[1].CreatedAt)
	require.NotNil(articles[1].UpdatedAt)

	// Check the content of the associated article files
	content, err := os.ReadFile("./testing/dirty/articles/article-1.md")
	require.Nil(err)
	require.Equal("# Article 1\n\n", string(content))

	content, err = os.ReadFile("./testing/dirty/articles/article-2.md")
	require.Nil(err)
	require.Equal("# Article 2\n\n", string(content))
}

// Test deleting an article
func TestDeleteArticle(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new article
	require.Nil(manager.CreateArticle("article-1"))

	// Get the article to verify it's been saved
	article := manager.GetArticle("article-1")
	require.NotNil(article)

	require.Equal("# Article 1\n\n", article.Content)
	require.Equal("article-1", article.Filename)
	require.Equal("Ethan", article.Author)
	require.NotNil(article.CreatedAt)
	require.NotNil(article.UpdatedAt)

	// Check the content of the associated article file
	content, err := os.ReadFile("./testing/dirty/articles/article-1.md")
	require.Nil(err)
	require.Equal("# Article 1\n\n", string(content))

	// Delete the article
	require.Nil(manager.DeleteArticle("article-1"))

	// Verify the article is gone
	require.Nil(manager.GetArticle("article-1"))

	// Verify the article file is gone
	_, err = os.Stat("./testing/dirty/articles/article-1.md")
	require.True(errors.Is(err, os.ErrNotExist))
}

// Test deleting an article that doesn't exist
func TestDeleteArticleNotFound(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Delete the article that doesn't exist
	err = manager.DeleteArticle("article-1")
	require.NotNil(err)
	require.Equal("article with name 'article-1' not found", err.Error())
}

// Test opening an article
func TestOpenArticle(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Create a new article
	require.Nil(manager.CreateArticle("article-1"))

	// Open the article
	require.Nil(manager.OpenArticle("article-1"))
}

// Test opening an article that doesn't exist
func TestOpenArticleNotFound(t *testing.T) {
	// Setup test
	require := require.New(t)
	manager, err := managerTestSetup()
	require.Nil(err)

	// Delete the article that doesn't exist
	err = manager.OpenArticle("article-1")
	require.NotNil(err)
	require.Equal("article with name 'article-1' not found", err.Error())
}

// Test loading a manager that already exists
func TestLoad(t *testing.T) {
	require := require.New(t)

	// Get working directory
	wd, err := os.Getwd()
	require.Nil(err)

	// Modify default paths to test files
	cli.ModifyDefaultDirectoryPath(path.Join(wd, "testing/pristine/articles/"))
	cli.ModifyConfigPath(path.Join(wd, "testing/pristine/config.json"))
	cli.ModifyManagerPath(path.Join(wd, "testing/pristine/manager.json"))

	// Get the default manager
	manager, err := cli.GetManager()
	require.Nil(err)

	// Check manager config
	require.Equal("testing/pristine/articles/", manager.Config.Directory)
	require.Equal("vi", manager.Config.Editor)
	require.Equal("John", manager.Config.DefaultAuthor)

	// Check article list
	require.Equal(2, len(manager.Articles))

	article1 := manager.Articles[0]
	require.Equal("article-1", article1.Filename)
	require.Equal("John", article1.Author)
	require.NotNil(article1.CreatedAt)
	require.NotNil(article1.UpdatedAt)

	article2 := manager.Articles[1]
	require.Equal("article-2", article2.Filename)
	require.Equal("Mary", article2.Author)
	require.NotNil(article2.CreatedAt)
	require.NotNil(article2.UpdatedAt)
}
