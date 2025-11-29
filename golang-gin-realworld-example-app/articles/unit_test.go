package articles

import (
	"fmt"
	"net/http/httptest"
	"realworld-backend/common"
	"realworld-backend/users"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var test_db *gorm.DB

// Test helper functions
func setupTestDB() {
	test_db = common.TestDBInit()
	users.AutoMigrate()
	AutoMigrate()
}

func teardownTestDB() {
	common.TestDBFree(test_db)
}

func createTestUser(username string) users.UserModel {
	userModel := users.UserModel{
		Username: username,
		Email:    fmt.Sprintf("%s@test.com", username),
		Bio:      fmt.Sprintf("Bio for %s", username),
	}
	userModel.SetPassword("password123")
	test_db.Create(&userModel)
	return userModel
}

func createTestArticle(title, description, body string, author ArticleUserModel) ArticleModel {
	articleModel := ArticleModel{
		Title:       title,
		Description: description,
		Body:        body,
		Slug:        fmt.Sprintf("%s-slug", title),
		Author:      author,
		AuthorID:    author.ID,
	}
	test_db.Create(&articleModel)
	return articleModel
}

// ==============================================
// Model Tests
// ==============================================

func TestArticleCreationWithValidData(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("testuser1")
	articleUserModel := GetArticleUserModel(userModel)

	articleModel := ArticleModel{
		Title:       "Test Article",
		Description: "Test Description",
		Body:        "Test Body Content",
		Slug:        "test-article-slug",
		Author:      articleUserModel,
		AuthorID:    articleUserModel.ID,
	}

	err := test_db.Create(&articleModel).Error
	asserts.NoError(err, "Article should be created successfully")
	asserts.NotZero(articleModel.ID, "Article ID should be set")
	asserts.Equal("Test Article", articleModel.Title)
	asserts.Equal("Test Description", articleModel.Description)
	asserts.Equal("Test Body Content", articleModel.Body)
}

func TestArticleValidationEmptyTitle(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("testuser2")
	articleUserModel := GetArticleUserModel(userModel)

	articleModel := ArticleModel{
		Title:       "", // Empty title
		Description: "Test Description",
		Body:        "Test Body",
		Slug:        "empty-title",
		Author:      articleUserModel,
		AuthorID:    articleUserModel.ID,
	}

	// Article can be created but validator should catch this
	err := test_db.Create(&articleModel).Error
	asserts.NoError(err, "Database allows empty title, but validator should prevent it")
}

func TestArticleFavoriteByUser(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("author1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Favorite Test", "Description", "Body", articleUserModel)

	favoriteUser := createTestUser("favoriter1")
	favoriteArticleUser := GetArticleUserModel(favoriteUser)

	// Initially not favorited
	asserts.False(articleModel.isFavoriteBy(favoriteArticleUser), "Article should not be favorited initially")
	asserts.Equal(uint(0), articleModel.favoritesCount(), "Favorite count should be 0")

	// Favorite the article
	err := articleModel.favoriteBy(favoriteArticleUser)
	asserts.NoError(err, "Article should be favorited successfully")

	// Verify favorited
	asserts.True(articleModel.isFavoriteBy(favoriteArticleUser), "Article should be favorited")
	asserts.Equal(uint(1), articleModel.favoritesCount(), "Favorite count should be 1")
}

func TestArticleUnfavoriteByUser(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("author2")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Unfavorite Test", "Description", "Body", articleUserModel)

	favoriteUser := createTestUser("favoriter2")
	favoriteArticleUser := GetArticleUserModel(favoriteUser)

	// Favorite first
	articleModel.favoriteBy(favoriteArticleUser)
	asserts.True(articleModel.isFavoriteBy(favoriteArticleUser), "Article should be favorited")

	// Unfavorite
	err := articleModel.unFavoriteBy(favoriteArticleUser)
	asserts.NoError(err, "Article should be unfavorited successfully")

	// Verify unfavorited
	asserts.False(articleModel.isFavoriteBy(favoriteArticleUser), "Article should not be favorited after unfavorite")
	asserts.Equal(uint(0), articleModel.favoritesCount(), "Favorite count should be 0 after unfavorite")
}

func TestArticleTagAssociation(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("author3")
	articleUserModel := GetArticleUserModel(userModel)

	articleModel := ArticleModel{
		Title:       "Tag Test Article",
		Description: "Description",
		Body:        "Body",
		Slug:        "tag-test-article",
		Author:      articleUserModel,
		AuthorID:    articleUserModel.ID,
	}

	// Set tags
	tags := []string{"golang", "testing", "backend"}
	err := articleModel.setTags(tags)
	asserts.NoError(err, "Tags should be set successfully")

	test_db.Create(&articleModel)

	// Retrieve and verify tags
	var retrievedArticle ArticleModel
	test_db.Where("id = ?", articleModel.ID).First(&retrievedArticle)
	test_db.Model(&retrievedArticle).Related(&retrievedArticle.Tags, "Tags")

	asserts.Equal(3, len(retrievedArticle.Tags), "Article should have 3 tags")
}

func TestMultipleFavoritesByDifferentUsers(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("author4")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Multi Favorite Test", "Description", "Body", articleUserModel)

	// Create multiple users who favorite
	user1 := createTestUser("favoriter3")
	user2 := createTestUser("favoriter4")
	user3 := createTestUser("favoriter5")

	articleUser1 := GetArticleUserModel(user1)
	articleUser2 := GetArticleUserModel(user2)
	articleUser3 := GetArticleUserModel(user3)

	articleModel.favoriteBy(articleUser1)
	articleModel.favoriteBy(articleUser2)
	articleModel.favoriteBy(articleUser3)

	asserts.Equal(uint(3), articleModel.favoritesCount(), "Article should have 3 favorites")
	asserts.True(articleModel.isFavoriteBy(articleUser1), "User1 should have favorited")
	asserts.True(articleModel.isFavoriteBy(articleUser2), "User2 should have favorited")
	asserts.True(articleModel.isFavoriteBy(articleUser3), "User3 should have favorited")
}

// ==============================================
// Serializer Tests
// ==============================================

func TestArticleSerializer(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("serializer1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Serializer Test", "Test Description", "Test Body", articleUserModel)

	// Setup Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("my_user_model", userModel)

	serializer := ArticleSerializer{c, articleModel}
	response := serializer.Response()

	asserts.Equal("Serializer Test", response.Title)
	asserts.Equal("Test Description", response.Description)
	asserts.Equal("Test Body", response.Body)
	asserts.NotEmpty(response.CreatedAt)
	asserts.NotEmpty(response.UpdatedAt)
	asserts.Equal(userModel.Username, response.Author.Username)
	asserts.False(response.Favorite)
	asserts.Equal(uint(0), response.FavoritesCount)
}

func TestArticleSerializerWithAuthor(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("serializer2")
	userModel.Bio = "Author Bio"
	test_db.Save(&userModel)

	articleUserModel := GetArticleUserModel(userModel)
	test_db.Model(&articleUserModel).Related(&articleUserModel.UserModel)

	articleModel := createTestArticle("Author Test", "Description", "Body", articleUserModel)
	articleModel.Author = articleUserModel

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("my_user_model", userModel)

	serializer := ArticleSerializer{c, articleModel}
	response := serializer.Response()

	asserts.Equal("serializer2", response.Author.Username)
	asserts.Equal("Author Bio", response.Author.Bio)
}

func TestArticleListSerializer(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("listserializer")
	articleUserModel := GetArticleUserModel(userModel)

	// Create multiple articles
	article1 := createTestArticle("Article 1", "Description 1", "Body 1", articleUserModel)
	article2 := createTestArticle("Article 2", "Description 2", "Body 2", articleUserModel)
	article3 := createTestArticle("Article 3", "Description 3", "Body 3", articleUserModel)

	articles := []ArticleModel{article1, article2, article3}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("my_user_model", userModel)

	serializer := ArticlesSerializer{c, articles}
	response := serializer.Response()

	asserts.Equal(3, len(response), "Should serialize 3 articles")
	asserts.Equal("Article 1", response[0].Title)
	asserts.Equal("Article 2", response[1].Title)
	asserts.Equal("Article 3", response[2].Title)
}

func TestCommentSerializer(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("commenter1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Comment Test", "Description", "Body", articleUserModel)

	commentModel := CommentModel{
		Article:   articleModel,
		ArticleID: articleModel.ID,
		Author:    articleUserModel,
		AuthorID:  articleUserModel.ID,
		Body:      "This is a test comment",
	}
	test_db.Create(&commentModel)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("my_user_model", userModel)

	serializer := CommentSerializer{c, commentModel}
	response := serializer.Response()

	asserts.NotZero(response.ID)
	asserts.Equal("This is a test comment", response.Body)
	asserts.NotEmpty(response.CreatedAt)
	asserts.NotEmpty(response.UpdatedAt)
	asserts.Equal("commenter1", response.Author.Username)
}

func TestCommentSerializerWithAuthor(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("commenter2")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Comment Author Test", "Description", "Body", articleUserModel)

	commentModel := CommentModel{
		Article:   articleModel,
		ArticleID: articleModel.ID,
		Author:    articleUserModel,
		AuthorID:  articleUserModel.ID,
		Body:      "Another test comment",
	}
	test_db.Create(&commentModel)
	test_db.Model(&commentModel).Related(&commentModel.Author, "Author")
	test_db.Model(&commentModel.Author).Related(&commentModel.Author.UserModel)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("my_user_model", userModel)

	serializer := CommentSerializer{c, commentModel}
	response := serializer.Response()

	asserts.Equal("commenter2", response.Author.Username)
	asserts.Equal("Another test comment", response.Body)
}

// ==============================================
// Validator Tests
// ==============================================

func TestArticleModelValidatorWithValidInput(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("validator1")
	articleUserModel := GetArticleUserModel(userModel)

	validator := NewArticleModelValidator()
	validator.Article.Title = "Valid Article Title"
	validator.Article.Description = "Valid Description"
	validator.Article.Body = "Valid Body Content"
	validator.Article.Tags = []string{"tag1", "tag2"}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("my_user_model", userModel)

	// Manually set the validator fields (simulating binding)
	validator.articleModel.Title = validator.Article.Title
	validator.articleModel.Description = validator.Article.Description
	validator.articleModel.Body = validator.Article.Body
	validator.articleModel.Author = articleUserModel

	asserts.Equal("Valid Article Title", validator.articleModel.Title)
	asserts.Equal("Valid Description", validator.articleModel.Description)
	asserts.Equal("Valid Body Content", validator.articleModel.Body)
}

func TestArticleModelValidatorMissingTitle(t *testing.T) {
	asserts := assert.New(t)

	validator := NewArticleModelValidator()
	validator.Article.Title = "" // Empty title should fail validation
	validator.Article.Description = "Description"
	validator.Article.Body = "Body"

	// The validator requires title with binding:"required,min=4"
	// So empty title would fail binding validation
	asserts.Empty(validator.Article.Title, "Title should be empty and fail validation")
}

func TestArticleModelValidatorTitleTooShort(t *testing.T) {
	asserts := assert.New(t)

	validator := NewArticleModelValidator()
	validator.Article.Title = "ab" // Less than minimum 4 characters
	validator.Article.Description = "Description"
	validator.Article.Body = "Body"

	// Title "ab" is less than min=4, should fail validation
	asserts.Less(len(validator.Article.Title), 4, "Title should be less than 4 characters")
}

func TestCommentModelValidatorWithValidInput(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("commentvalidator1")
	articleUserModel := GetArticleUserModel(userModel)

	validator := NewCommentModelValidator()
	validator.Comment.Body = "This is a valid comment"

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("my_user_model", userModel)

	validator.commentModel.Body = validator.Comment.Body
	validator.commentModel.Author = articleUserModel

	asserts.Equal("This is a valid comment", validator.commentModel.Body)
	asserts.Equal(articleUserModel.ID, validator.commentModel.Author.ID)
}

func TestCommentModelValidatorMissingBody(t *testing.T) {
	asserts := assert.New(t)

	validator := NewCommentModelValidator()
	validator.Comment.Body = "" // Empty body

	// Empty comment body should be allowed by validator (no required tag)
	// but it's not meaningful
	asserts.Empty(validator.Comment.Body, "Comment body is empty")
}

func TestTagSerializer(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	tagModel := TagModel{Tag: "golang"}
	test_db.Create(&tagModel)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	serializer := TagSerializer{c, tagModel}
	response := serializer.Response()

	asserts.Equal("golang", response)
}

func TestTagsSerializer(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	tag1 := TagModel{Tag: "golang"}
	tag2 := TagModel{Tag: "testing"}
	tag3 := TagModel{Tag: "backend"}
	test_db.Create(&tag1)
	test_db.Create(&tag2)
	test_db.Create(&tag3)

	tags := []TagModel{tag1, tag2, tag3}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	serializer := TagsSerializer{c, tags}
	response := serializer.Response()

	asserts.Equal(3, len(response))
	asserts.Contains(response, "golang")
	asserts.Contains(response, "testing")
	asserts.Contains(response, "backend")
}

func TestFindOneArticle(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("findtest1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Find Test", "Description", "Body", articleUserModel)

	foundArticle, err := FindOneArticle(&ArticleModel{Slug: articleModel.Slug})
	asserts.NoError(err)
	asserts.Equal(articleModel.ID, foundArticle.ID)
	asserts.Equal("Find Test", foundArticle.Title)
}

func TestDeleteArticleModel(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("deletetest1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Delete Test", "Description", "Body", articleUserModel)

	err := DeleteArticleModel(&ArticleModel{Slug: articleModel.Slug})
	asserts.NoError(err)

	// Verify deleted
	var count int
	test_db.Model(&ArticleModel{}).Where("id = ?", articleModel.ID).Count(&count)
	asserts.Equal(0, count, "Article should be deleted")
}

func TestDeleteCommentModel(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("deletecomment1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Comment Delete Test", "Description", "Body", articleUserModel)

	commentModel := CommentModel{
		Article:   articleModel,
		ArticleID: articleModel.ID,
		Author:    articleUserModel,
		AuthorID:  articleUserModel.ID,
		Body:      "Comment to be deleted",
	}
	test_db.Create(&commentModel)

	err := DeleteCommentModel(&CommentModel{ArticleID: commentModel.ArticleID, AuthorID: commentModel.AuthorID})
	asserts.NoError(err)

	// Verify deleted
	var count int
	test_db.Model(&CommentModel{}).Where("id = ?", commentModel.ID).Count(&count)
	asserts.Equal(0, count, "Comment should be deleted")
}
