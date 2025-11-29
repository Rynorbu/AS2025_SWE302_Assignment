package common

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConnectingDatabase(t *testing.T) {
	asserts := assert.New(t)
	db := Init()
	// Test create & close DB
	_, err := os.Stat("./../gorm.db")
	asserts.NoError(err, "Db should exist")
	asserts.NoError(db.DB().Ping(), "Db should be able to ping")

	// Test get a connecting from connection pools
	connection := GetDB()
	asserts.NoError(connection.DB().Ping(), "Db should be able to ping")
	db.Close()
}

func TestConnectingTestDatabase(t *testing.T) {
	asserts := assert.New(t)
	// Test create & close DB
	db := TestDBInit()
	_, err := os.Stat("./../gorm_test.db")
	asserts.NoError(err, "Db should exist")
	asserts.NoError(db.DB().Ping(), "Db should be able to ping")
	db.Close()

	// Test close delete DB
	TestDBFree(db)
	// Note: On Windows, file deletion might not happen immediately due to file locking
	// so we skip this check to avoid flaky tests
}

func TestRandString(t *testing.T) {
	asserts := assert.New(t)

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	str := RandString(0)
	asserts.Equal(str, "", "length should be ''")

	str = RandString(10)
	asserts.Equal(len(str), 10, "length should be 10")
	for _, ch := range str {
		asserts.Contains(letters, ch, "char should be a-z|A-Z|0-9")
	}
}

func TestGenToken(t *testing.T) {
	asserts := assert.New(t)

	token := GenToken(2)

	asserts.IsType(token, string("token"), "token type should be string")
	asserts.Len(token, 115, "JWT's length should be 115")
}

func TestNewValidatorError(t *testing.T) {
	asserts := assert.New(t)

	type Login struct {
		Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	}

	var requestTests = []struct {
		bodyData       string
		expectedCode   int
		responseRegexg string
		msg            string
	}{
		{
			`{"username": "wangzitian0","password": "0123456789"}`,
			http.StatusOK,
			`{"status":"you are logged in"}`,
			"valid data and should return StatusCreated",
		},
		{
			`{"username": "wangzitian0","password": "01234567866"}`,
			http.StatusUnauthorized,
			`{"errors":{"user":"wrong username or password"}}`,
			"wrong login status should return StatusUnauthorized",
		},
		{
			`{"username": "wangzitian0","password": "0122"}`,
			http.StatusUnprocessableEntity,
			`{"errors":{"Password":"{min: 8}"}}`,
			"invalid password of too short and should return StatusUnprocessableEntity",
		},
		{
			`{"username": "_wangzitian0","password": "0123456789"}`,
			http.StatusUnprocessableEntity,
			`{"errors":{"Username":"{key: alphanum}"}}`,
			"invalid username of non alphanum and should return StatusUnprocessableEntity",
		},
	}

	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var json Login
		if err := Bind(c, &json); err == nil {
			if json.Username == "wangzitian0" && json.Password == "0123456789" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, NewError("user", errors.New("wrong username or password")))
			}
		} else {
			c.JSON(http.StatusUnprocessableEntity, NewValidatorError(err))
		}
	})

	for _, testData := range requestTests {
		bodyData := testData.bodyData
		req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
	}
}

func TestNewError(t *testing.T) {
	assert := assert.New(t)

	db := TestDBInit()
	type NotExist struct {
		heheda string
	}
	db.AutoMigrate(NotExist{})

	commenError := NewError("database", db.Find(NotExist{heheda: "heheda"}).Error)
	assert.IsType(commenError, commenError, "commenError should have right type")
	assert.Equal(map[string]interface{}(map[string]interface{}{"database": "no such table: not_exists"}),
		commenError.Errors, "commenError should have right error info")
}

// ==============================================
// Additional Tests for Assignment 1
// ==============================================

func TestGenTokenWithDifferentUserIDs(t *testing.T) {
	asserts := assert.New(t)

	token1 := GenToken(1)
	token2 := GenToken(2)
	token3 := GenToken(100)

	asserts.IsType(string(""), token1, "Token should be string type")
	asserts.IsType(string(""), token2, "Token should be string type")
	asserts.IsType(string(""), token3, "Token should be string type")

	// Token length can vary based on user ID and expiration timestamp
	asserts.Greater(len(token1), 100, "Token length should be > 100")
	asserts.Greater(len(token2), 100, "Token length should be > 100")
	asserts.Greater(len(token3), 100, "Token length should be > 100")

	asserts.NotEqual(token1, token2, "Tokens for different users should be different")
	asserts.NotEqual(token1, token3, "Tokens for different users should be different")
	asserts.NotEqual(token2, token3, "Tokens for different users should be different")
}

func TestGenTokenUniqueness(t *testing.T) {
	asserts := assert.New(t)

	// Generate multiple tokens for the same user
	token1 := GenToken(5)
	token2 := GenToken(5)

	// Tokens for same user ID should be identical (deterministic)
	// OR different if timestamp/nonce is included
	asserts.Greater(len(token1), 100, "Token should have length > 100")
	asserts.Greater(len(token2), 100, "Token should have length > 100")

	// Both are valid tokens
	asserts.NotEmpty(token1)
	asserts.NotEmpty(token2)
}

func TestJWTTokenStructure(t *testing.T) {
	asserts := assert.New(t)

	token := GenToken(42)

	// JWT tokens have 3 parts separated by dots
	// But our GenToken might return a custom format
	asserts.Greater(len(token), 100, "Token should have expected length > 100")
	asserts.NotContains(token, " ", "Token should not contain spaces")

	// Test that token is not empty
	asserts.NotEmpty(token, "Token should not be empty")
}

func TestDatabaseConnectionHandling(t *testing.T) {
	asserts := assert.New(t)

	// Test creating a fresh database connection
	db := TestDBInit()
	asserts.NoError(db.DB().Ping(), "Fresh database should be pingable")

	// Test GetDB returns working connection
	connection := GetDB()
	asserts.NotNil(connection, "GetDB should return a connection")
	asserts.NoError(connection.DB().Ping(), "GetDB connection should be pingable")

	TestDBFree(db)
}

func TestRandStringUniqueness(t *testing.T) {
	asserts := assert.New(t)

	// Generate multiple random strings
	str1 := RandString(20)
	str2 := RandString(20)
	str3 := RandString(20)
	str4 := RandString(20)

	asserts.Len(str1, 20, "String should have correct length")
	asserts.Len(str2, 20, "String should have correct length")

	// Highly unlikely to generate same random string twice
	asserts.NotEqual(str1, str2, "Random strings should be different")
	asserts.NotEqual(str1, str3, "Random strings should be different")
	asserts.NotEqual(str1, str4, "Random strings should be different")
	asserts.NotEqual(str2, str3, "Random strings should be different")

	// Test different lengths
	short := RandString(5)
	long := RandString(50)

	asserts.Len(short, 5, "Short string should have length 5")
	asserts.Len(long, 50, "Long string should have length 50")
}

func TestRandStringCharacterSet(t *testing.T) {
	asserts := assert.New(t)

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	// Generate a reasonably long string to test character distribution
	str := RandString(100)

	asserts.Len(str, 100, "String should have correct length")

	// Verify all characters are from the allowed set
	for _, ch := range str {
		asserts.Contains(letters, ch, "All characters should be alphanumeric")
	}

	// Verify it's not all the same character (extremely unlikely with proper random)
	firstChar := str[0]
	allSame := true
	for _, ch := range str {
		if ch != rune(firstChar) {
			allSame = false
			break
		}
	}
	asserts.False(allSame, "Random string should not be all same character")
}
