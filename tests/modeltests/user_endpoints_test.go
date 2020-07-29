package modeltests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/garcialuis/Gossip/api/models"
	"gopkg.in/go-playground/assert.v1"
)

var jwtToken = models.Token{}
var createdUserID uint32

func TestHomeEndpoint(t *testing.T) {

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	server.Router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

}

// TestGetusers : returns all users (2) seeded in the database
func TestGetUsers(t *testing.T) {

	var users []models.User

	// Call refreshUserTable to clear the user table:
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	// seeUsers will then populate the table with two users:
	err = seedUsers()
	if err != nil {
		log.Fatal("Could not seed users: ", err)
	}

	// Prepare the request for the endpoint that's being tested:
	request, _ := http.NewRequest("GET", "/users", nil)
	// The Recorder will be used to compare the response values later:
	response := httptest.NewRecorder()

	server.Router.ServeHTTP(response, request)

	// Convert the response body into a []byte, then unmarshal:
	responseBody, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(responseBody, &users)

	// Use assertions to compare results to expected values:
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, 2, len(users))

	assert.Equal(t, "Lucky", users[0].Nickname)
	assert.Equal(t, "lucky@email.com", users[0].Email)

	assert.Equal(t, "Sope", users[1].Nickname)
	assert.Equal(t, "sope@email.com", users[1].Email)
}

// TestCreateUser : Ensures that a user can be created, sets createdUserID that's used by other tests.
func TestCreateUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	newUser := models.User{
		Nickname: "Tommy T",
		Email:    "tommyt@email.com",
		Password: "dummypassword",
	}

	// Convert newUser struct to io.Reader
	newUserByte, _ := json.Marshal(newUser)
	requestReader := bytes.NewReader(newUserByte)

	request, _ := http.NewRequest("POST", "/users", requestReader)
	response := httptest.NewRecorder()

	server.Router.ServeHTTP(response, request)

	responseBody, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(responseBody, &newUser)

	// Set userID from created user, will be needed for other tests
	createdUserID = newUser.ID

	assert.Equal(t, 201, response.Code)
}

// TestLoginUser : logs in the user that was created in the previous test,
//				also, sets jwtToken required in authenticated requests
func TestLoginUser(t *testing.T) {

	loginBody := models.User{
		Email:    "tommyt@email.com",
		Password: "dummypassword",
	}

	userBytes, _ := json.Marshal(loginBody)
	bodyIOReader := bytes.NewReader(userBytes)

	request, _ := http.NewRequest("POST", "/login", bodyIOReader)
	response := httptest.NewRecorder()

	server.Router.ServeHTTP(response, request)

	// Response includes jwt token:
	responseBody, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(responseBody, &jwtToken)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, response.Code)
}

// TestGetUser : tests authenticated /users endpoint.
func TestGetUser(t *testing.T) {

	url := fmt.Sprint("/private/users/", createdUserID)
	authToken := "Bearer " + jwtToken.Token

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", authToken)
	response := httptest.NewRecorder()

	server.Router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

// GetValidToken : Used by tests making requests to authenticated endpoints.
// 	It checks to see if the current jwtToken is valid, if it's not it will login the user
//	to create and returns the new valid token.
func GetValidToken() string {

	url := fmt.Sprint("/private/users/", createdUserID)
	authToken := "Bearer " + jwtToken.Token

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", authToken)
	response := httptest.NewRecorder()

	server.Router.ServeHTTP(response, request)

	if response.Code == 401 {
		fmt.Print("GOT A 401 BACK SO WE WILL LOG BACK IN")
		login()
	}

	return jwtToken.Token
}

// login : function used to login user to generate a jwtToken
func login() {

	loginBody := models.User{
		Email:    "tommyt@email.com",
		Password: "dummypassword",
	}

	userBytes, _ := json.Marshal(loginBody)
	bodyIOReader := bytes.NewReader(userBytes)

	request, _ := http.NewRequest("POST", "/login", bodyIOReader)
	response := httptest.NewRecorder()

	server.Router.ServeHTTP(response, request)

	responseBody, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(responseBody, &jwtToken)

	if err != nil {
		log.Fatal(err)
	}

}
