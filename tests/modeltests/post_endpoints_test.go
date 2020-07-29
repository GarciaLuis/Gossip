package modeltests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetPostsEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/posts", nil)
	response := httptest.NewRecorder()

	server.Router.ServeHTTP(response, request)

	// /posts
	fmt.Println("THIS IS THE RETURN STATUS CODE: ", response.Code)
	fmt.Println("THIS IS THE RESPONSE: ", response.Body.String())

	assert.Equal(t, 200, response.Code)
}
