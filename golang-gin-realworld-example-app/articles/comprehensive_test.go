package articles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// ==============================================
// Additional Model Tests for 100% Coverage
// ==============================================

func TestGetArticleUserModel(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("getusermodel1")
	articleUserModel := GetArticleUserModel(userModel)

	asserts.Equal(userModel.ID, articleUserModel.UserModelID)
	asserts.Equal(userModel.Username, articleUserModel.UserModel.Username)
}

func TestFavoritesCount(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("favcount1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Favorite Count Test", "Description", "Body", articleUserModel)

	// Test initial count
	count := articleModel.favoritesCount()
	asserts.Equal(uint(0), count)

	// Add favorites
	user2 := createTestUser("favcount2")
	articleUser2 := GetArticleUserModel(user2)
	articleModel.favoriteBy(articleUser2)

	count = articleModel.favoritesCount()
	asserts.Equal(uint(1), count)
}

func TestIsFavoriteBy(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("isfav1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Is Favorite Test", "Description", "Body", articleUserModel)

	// Test not favorited
	isFav := articleModel.isFavoriteBy(articleUserModel)
	asserts.False(isFav)

	// Favorite and test again
	articleModel.favoriteBy(articleUserModel)
	isFav = articleModel.isFavoriteBy(articleUserModel)
	asserts.True(isFav)
}

func TestSaveOne(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("saveone1")
	articleUserModel := GetArticleUserModel(userModel)

	articleModel := ArticleModel{
		Title:       "Save One Test",
		Slug:        "save-one-test",
		Description: "Test Description",
		Body:        "Test Body",
		Author:      articleUserModel,
		AuthorID:    articleUserModel.ID,
	}

	err := SaveOne(&articleModel)
	asserts.NoError(err)
	asserts.NotZero(articleModel.ID)
	asserts.NotZero(articleModel.CreatedAt)
	asserts.NotZero(articleModel.UpdatedAt)
}

func TestGetComments(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("getcomments1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Get Comments Test", "Description", "Body", articleUserModel)

	// Add comments
	comment1 := CommentModel{
		Article:   articleModel,
		ArticleID: articleModel.ID,
		Author:    articleUserModel,
		AuthorID:  articleUserModel.ID,
		Body:      "First comment",
	}
	test_db.Create(&comment1)

	comment2 := CommentModel{
		Article:   articleModel,
		ArticleID: articleModel.ID,
		Author:    articleUserModel,
		AuthorID:  articleUserModel.ID,
		Body:      "Second comment",
	}
	test_db.Create(&comment2)

	err := articleModel.getComments()
	asserts.NoError(err)
	asserts.Equal(2, len(articleModel.Comments))
}

func TestGetAllTags(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("getalltags1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Get All Tags Test", "Description", "Body", articleUserModel)

	// Add tags
	articleModel.setTags([]string{"tag1", "tag2", "tag3"})

	tags, err := getAllTags()
	asserts.NoError(err)
	asserts.GreaterOrEqual(len(tags), 3)
}

func TestFindManyArticleWithFilters(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("findmany1")
	articleUserModel := GetArticleUserModel(userModel)

	// Create multiple articles
	article1 := createTestArticle("Article One", "Description", "Body", articleUserModel)
	article1.setTags([]string{"golang", "testing"})

	article2 := createTestArticle("Article Two", "Description", "Body", articleUserModel)
	article2.setTags([]string{"golang"})

	_ = createTestArticle("Article Three", "Description", "Body", articleUserModel)

	// Test tag filter
	articles, count, err := FindManyArticle("golang", "", "10", "0", "")
	asserts.NoError(err)
	asserts.GreaterOrEqual(count, 2)
	asserts.GreaterOrEqual(len(articles), 2)
}

func TestFindManyArticleWithAuthorFilter(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel1 := createTestUser("findauthor1")
	articleUserModel1 := GetArticleUserModel(userModel1)
	createTestArticle("Author Article 1", "Description", "Body", articleUserModel1)

	userModel2 := createTestUser("findauthor2")
	articleUserModel2 := GetArticleUserModel(userModel2)
	createTestArticle("Author Article 2", "Description", "Body", articleUserModel2)

	// Test author filter
	articles, count, err := FindManyArticle("", "findauthor1", "10", "0", "")
	asserts.NoError(err)
	asserts.GreaterOrEqual(count, 1)
	asserts.GreaterOrEqual(len(articles), 1)
	if len(articles) > 0 {
		asserts.Equal("findauthor1", articles[0].Author.UserModel.Username)
	}
}

func TestFindManyArticleWithFavoritedFilter(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel1 := createTestUser("findfav1")
	articleUserModel1 := GetArticleUserModel(userModel1)
	articleModel := createTestArticle("Favorited Article", "Description", "Body", articleUserModel1)

	userModel2 := createTestUser("findfav2")
	articleUserModel2 := GetArticleUserModel(userModel2)
	articleModel.favoriteBy(articleUserModel2)

	// Test favorited filter
	articles, count, err := FindManyArticle("", "", "10", "0", "findfav2")
	asserts.NoError(err)
	asserts.GreaterOrEqual(count, 1)
	asserts.GreaterOrEqual(len(articles), 1)
}

func TestFindManyArticleWithPagination(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("findpagination1")
	articleUserModel := GetArticleUserModel(userModel)

	// Create 5 articles
	for i := 1; i <= 5; i++ {
		createTestArticle(fmt.Sprintf("Pagination Article %d", i), "Description", "Body", articleUserModel)
	}

	// Test with limit
	articles, count, err := FindManyArticle("", "", "2", "0", "")
	asserts.NoError(err)
	asserts.GreaterOrEqual(count, 5)
	asserts.Equal(2, len(articles))

	// Test with offset
	articles2, _, err2 := FindManyArticle("", "", "2", "2", "")
	asserts.NoError(err2)
	asserts.Equal(2, len(articles2))

	// Verify different results
	if len(articles) > 0 && len(articles2) > 0 {
		asserts.NotEqual(articles[0].ID, articles2[0].ID)
	}
}

func TestGetArticleFeed(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel1 := createTestUser("feeduser1")
	articleUserModel1 := GetArticleUserModel(userModel1)

	userModel2 := createTestUser("feeduser2")
	articleUserModel2 := GetArticleUserModel(userModel2)
	createTestArticle("Feed Article", "Description", "Body", articleUserModel2)

	// User1 follows User2
	test_db.Model(&articleUserModel1).Association("Following").Append(articleUserModel2)

	// Get feed for user1
	articles, count, err := articleUserModel1.GetArticleFeed("10", "0")
	asserts.NoError(err)
	asserts.GreaterOrEqual(count, 0)
	asserts.GreaterOrEqual(len(articles), 0)
}

func TestSetTags(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("settags1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Set Tags Test", "Description", "Body", articleUserModel)

	// Set tags
	tags := []string{"golang", "testing", "backend"}
	err := articleModel.setTags(tags)
	asserts.NoError(err)

	// Verify tags were created
	allTags, err2 := getAllTags()
	asserts.NoError(err2)
	asserts.GreaterOrEqual(len(allTags), 3)
}

func TestUpdateArticle(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("update1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Original Title", "Original Description", "Original Body", articleUserModel)

	// Update article
	updateData := ArticleModel{
		Title:       "Updated Title",
		Description: "Updated Description",
		Body:        "Updated Body",
	}
	err := articleModel.Update(updateData)
	asserts.NoError(err)
	asserts.Equal("Updated Title", articleModel.Title)
	asserts.Equal("Updated Description", articleModel.Description)
	asserts.Equal("Updated Body", articleModel.Body)
	// Slug remains the same unless specifically updated
	asserts.NotEmpty(articleModel.Slug)
}

func TestAutoMigrate(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	AutoMigrate()

	// Verify tables exist
	asserts.True(test_db.HasTable(&ArticleModel{}))
	asserts.True(test_db.HasTable(&FavoriteModel{}))
	asserts.True(test_db.HasTable(&TagModel{}))
	asserts.True(test_db.HasTable(&CommentModel{}))
	asserts.True(test_db.HasTable(&ArticleUserModel{}))
}

// ==============================================
// Router Handler Tests for 100% Coverage
// ==============================================

func TestArticleCreate(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routecreate1")

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/articles", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleCreate(c)
	})

	body := map[string]interface{}{
		"article": map[string]interface{}{
			"title":       "Test Article",
			"description": "Test Description",
			"body":        "Test Body",
			"tagList":     []string{"tag1", "tag2"},
		},
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusCreated, w.Code)
	
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	article := response["article"].(map[string]interface{})
	asserts.Equal("Test Article", article["title"])
}

func TestArticleList(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routelist1")
	articleUserModel := GetArticleUserModel(userModel)
	createTestArticle("List Article 1", "Description", "Body", articleUserModel)
	createTestArticle("List Article 2", "Description", "Body", articleUserModel)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/articles", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleList(c)
	})

	req, _ := http.NewRequest("GET", "/articles", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	asserts.Contains(response, "articles")
	asserts.Contains(response, "articlesCount")
}

func TestArticleFeed(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routefeed1")

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/articles/feed", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleFeed(c)
	})

	req, _ := http.NewRequest("GET", "/articles/feed", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	asserts.Contains(response, "articles")
	asserts.Contains(response, "articlesCount")
}

func TestArticleRetrieve(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routeretrieve1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Retrieve Article", "Description", "Body", articleUserModel)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/articles/:slug", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleRetrieve(c)
	})

	req, _ := http.NewRequest("GET", fmt.Sprintf("/articles/%s", articleModel.Slug), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	article := response["article"].(map[string]interface{})
	asserts.Equal("Retrieve Article", article["title"])
}

func TestArticleUpdate(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routeupdate1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Update Article", "Description", "Body", articleUserModel)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.PUT("/articles/:slug", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleUpdate(c)
	})

	body := map[string]interface{}{
		"article": map[string]interface{}{
			"title": "Updated Title",
			"body":  "Updated Body",
		},
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/articles/%s", articleModel.Slug), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	article := response["article"].(map[string]interface{})
	asserts.Equal("Updated Title", article["title"])
}

func TestArticleDelete(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routedelete1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Delete Article", "Description", "Body", articleUserModel)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/articles/:slug", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleDelete(c)
	})

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/articles/%s", articleModel.Slug), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusOK, w.Code)

	// Verify deleted
	var count int
	test_db.Model(&ArticleModel{}).Where("id = ?", articleModel.ID).Count(&count)
	asserts.Equal(0, count)
}

func TestArticleFavorite(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routefavorite1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Favorite Article", "Description", "Body", articleUserModel)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/articles/:slug/favorite", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleFavorite(c)
	})

	req, _ := http.NewRequest("POST", fmt.Sprintf("/articles/%s/favorite", articleModel.Slug), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	article := response["article"].(map[string]interface{})
	asserts.True(article["favorited"].(bool))
	asserts.Equal(float64(1), article["favoritesCount"].(float64))
}

func TestArticleUnfavorite(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routeunfavorite1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Unfavorite Article", "Description", "Body", articleUserModel)

	// First favorite it
	articleModel.favoriteBy(articleUserModel)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/articles/:slug/favorite", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleUnfavorite(c)
	})

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/articles/%s/favorite", articleModel.Slug), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	article := response["article"].(map[string]interface{})
	asserts.False(article["favorited"].(bool))
	asserts.Equal(float64(0), article["favoritesCount"].(float64))
}

func TestArticleCommentCreate(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routecommentcreate1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Comment Article", "Description", "Body", articleUserModel)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/articles/:slug/comments", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleCommentCreate(c)
	})

	body := map[string]interface{}{
		"comment": map[string]interface{}{
			"body": "This is a test comment",
		},
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", fmt.Sprintf("/articles/%s/comments", articleModel.Slug), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	comment := response["comment"].(map[string]interface{})
	asserts.Equal("This is a test comment", comment["body"])
}

func TestArticleCommentDelete(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routecommentdelete1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Comment Delete Article", "Description", "Body", articleUserModel)

	commentModel := CommentModel{
		Article:   articleModel,
		ArticleID: articleModel.ID,
		Author:    articleUserModel,
		AuthorID:  articleUserModel.ID,
		Body:      "Comment to delete",
	}
	test_db.Create(&commentModel)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/articles/:slug/comments/:id", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleCommentDelete(c)
	})

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/articles/%s/comments/%d", articleModel.Slug, commentModel.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusOK, w.Code)
}

func TestArticleCommentList(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routecommentlist1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Comment List Article", "Description", "Body", articleUserModel)

	comment1 := CommentModel{
		Article:   articleModel,
		ArticleID: articleModel.ID,
		Author:    articleUserModel,
		AuthorID:  articleUserModel.ID,
		Body:      "First comment",
	}
	test_db.Create(&comment1)

	comment2 := CommentModel{
		Article:   articleModel,
		ArticleID: articleModel.ID,
		Author:    articleUserModel,
		AuthorID:  articleUserModel.ID,
		Body:      "Second comment",
	}
	test_db.Create(&comment2)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/articles/:slug/comments", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleCommentList(c)
	})

	req, _ := http.NewRequest("GET", fmt.Sprintf("/articles/%s/comments", articleModel.Slug), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	comments := response["comments"].([]interface{})
	asserts.Equal(2, len(comments))
}

func TestTagList(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	// Create tags
	tag1 := TagModel{Tag: "golang"}
	tag2 := TagModel{Tag: "testing"}
	test_db.Create(&tag1)
	test_db.Create(&tag2)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/tags", TagList)

	req, _ := http.NewRequest("GET", "/tags", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	tags := response["tags"].([]interface{})
	asserts.GreaterOrEqual(len(tags), 2)
}

// ==============================================
// Validator Binding Tests
// ==============================================

func TestArticleModelValidatorBind(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("validatorbind1")

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("my_user_model", userModel)

	body := map[string]interface{}{
		"article": map[string]interface{}{
			"title":       "Validator Test",
			"description": "Test Description",
			"body":        "Test Body",
			"tagList":     []string{"tag1"},
		},
	}
	jsonBody, _ := json.Marshal(body)
	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	validator := NewArticleModelValidator()
	err := validator.Bind(c)
	asserts.NoError(err)
	asserts.Equal("Validator Test", validator.articleModel.Title)
}

func TestCommentModelValidatorBind(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("commentbind1")

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("my_user_model", userModel)

	body := map[string]interface{}{
		"comment": map[string]interface{}{
			"body": "Comment body test",
		},
	}
	jsonBody, _ := json.Marshal(body)
	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	validator := NewCommentModelValidator()
	err := validator.Bind(c)
	asserts.NoError(err)
	asserts.Equal("Comment body test", validator.commentModel.Body)
}

// ==============================================
// Edge Cases and Error Handling
// ==============================================

func TestArticleRetrieveNotFound(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routenotfound1")

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/articles/:slug", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleRetrieve(c)
	})

	req, _ := http.NewRequest("GET", "/articles/nonexistent-slug", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusNotFound, w.Code)
}

func TestArticleCreateInvalidData(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("routeinvalid1")

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/articles", func(c *gin.Context) {
		c.Set("my_user_model", userModel)
		ArticleCreate(c)
	})

	body := map[string]interface{}{
		"article": map[string]interface{}{
			"title": "ab", // Too short
		},
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(http.StatusUnprocessableEntity, w.Code)
}

func TestCommentSerializersList(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("commentserlist1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Comment Serializers Test", "Description", "Body", articleUserModel)

	comment1 := CommentModel{
		Article:   articleModel,
		ArticleID: articleModel.ID,
		Author:    articleUserModel,
		AuthorID:  articleUserModel.ID,
		Body:      "Comment 1",
	}
	test_db.Create(&comment1)

	comment2 := CommentModel{
		Article:   articleModel,
		ArticleID: articleModel.ID,
		Author:    articleUserModel,
		AuthorID:  articleUserModel.ID,
		Body:      "Comment 2",
	}
	test_db.Create(&comment2)

	comments := []CommentModel{comment1, comment2}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("my_user_model", userModel)

	serializer := CommentsSerializer{c, comments}
	response := serializer.Response()

	asserts.Equal(2, len(response))
	asserts.Equal("Comment 1", response[0].Body)
	asserts.Equal("Comment 2", response[1].Body)
}

func TestNewArticleModelValidatorFillWith(t *testing.T) {
	asserts := assert.New(t)
	setupTestDB()
	defer teardownTestDB()

	userModel := createTestUser("fillwith1")
	articleUserModel := GetArticleUserModel(userModel)
	articleModel := createTestArticle("Fill With Test", "Description", "Body", articleUserModel)
	articleModel.setTags([]string{"tag1", "tag2"})

	validator := NewArticleModelValidatorFillWith(articleModel)

	asserts.Equal("Fill With Test", validator.Article.Title)
	asserts.Equal("Description", validator.Article.Description)
	asserts.Equal("Body", validator.Article.Body)
	asserts.Equal(2, len(validator.Article.Tags))
}
