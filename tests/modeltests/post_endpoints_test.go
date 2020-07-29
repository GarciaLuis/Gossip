package modeltests

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetPostsEndpoint(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and posts tables %v\n", err)
	}
	_, _, err = seedUsersAndPosts()
	if err != nil {
		log.Fatalf("Error seeding users and posts into table: %v\n", err)
	}

	request, _ := http.NewRequest("GET", "/posts", nil)
	response := httptest.NewRecorder()

	server.Router.ServeHTTP(response, request)

	// /posts
	fmt.Println("THIS IS THE RETURN STATUS CODE: ", response.Code)
	fmt.Println("THIS IS THE RESPONSE: ", response.Body.String())

	assert.Equal(t, 200, response.Code)
}
