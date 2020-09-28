package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/garcialuis/Gossip/api/models"
	"gopkg.in/go-playground/assert.v1"
)

var testServer = Server{}

func TestMain(m *testing.M) {

	testServer.InitializeTestServer(os.Getenv("DB_POSTGRES_DRIVER"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_NAME"))
	os.Exit(m.Run())
}

func TestGetPostsEndpoint(t *testing.T) {

	err := models.RefreshUserAndPostTable(testServer.DB)
	if err != nil {
		log.Fatalf("Error refreshing user and posts tables %v\n", err)
	}
	_, _, err = models.SeedUsersAndPosts(testServer.DB)
	if err != nil {
		log.Fatalf("Error seeding users and posts into table: %v\n", err)
	}

	request, _ := http.NewRequest("GET", "/posts", nil)
	response := httptest.NewRecorder()

	testServer.Router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

// TestGetPostEndpoint : Verifies that we can retrieve a post by its id.
func TestGetPostEnpoint(t *testing.T) {

	err := models.RefreshUserAndPostTable(testServer.DB)
	if err != nil {
		log.Fatalf("Error refreshing user and posts tables %v\n", err)
	}

	_, _, err = models.SeedUsersAndPosts(testServer.DB)
	if err != nil {
		log.Fatalf("Error seeding users and posts into table: %v\n", err)
	}

	request, _ := http.NewRequest("GET", "/posts/1", nil)
	response := httptest.NewRecorder()

	testServer.Router.ServeHTTP(response, request)

	post := models.Post{}

	responseBody, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(responseBody, &post)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Post title Num 1", post.Title)
	assert.Equal(t, "Content of Post Num 1", post.Content)
}

// TestUpdatePostEndpoint : ensures that a post can be updated by the author.
func TestUpdatePostEndpoint(t *testing.T) {

	err := models.RefreshUserAndPostTable(testServer.DB)
	if err != nil {
		log.Fatalf("Error refreshing user and posts tables %v\n", err)
	}

	users, _, err := models.SeedUsersAndPosts(testServer.DB)
	if err != nil {
		log.Fatalf("Error seeding users and posts into table: %v\n", err)
	}

	author := users[0]
	GetValidToken(author.Email, "dummypassword")

	updatedPost := models.Post{
		Title:    "update post title",
		Content:  "Content of update post",
		AuthorID: author.ID,
	}

	updatedPostBytes, err := json.Marshal(updatedPost)
	updatedPostBody := bytes.NewReader(updatedPostBytes)

	authToken := "Bearer " + jwtToken.Token
	request, _ := http.NewRequest("PUT", "/posts/1", updatedPostBody)
	request.Header.Add("Authorization", authToken)
	response := httptest.NewRecorder()

	testServer.Router.ServeHTTP(response, request)

	modifiedPost := models.Post{}
	resBody, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(resBody, &modifiedPost)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, updatedPost.Title, modifiedPost.Title)
	assert.Equal(t, updatedPost.Content, modifiedPost.Content)
}
