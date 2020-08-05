// Package classification Gossip API
//
// Documentation for Gossip API
//
//		Schemes: http
//		BasePath: /
//		title: Gossip API
//		version: 1.0.0
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//
//		Security:
//		- api_key:
//
//		SecurityDefinitions:
//		api_key:
//			type: apiKey
//			name: KEY
//			in: header
//
// swagger:meta

package controllers

import "github.com/garcialuis/Gossip/api/models"

// MODELS USED TO DOCUMENT USER ENDPOINTS:

// User record
// swagger:response userResponse
type userResponseWrapper struct {
	// Single user record
	// in: body
	Body models.User
}

// User record that may include sensitive information
// swagger:response authenticatedUser
type authenticatedUserResponse struct {
	// Single user record
	// in: body
	Body models.User
}

// A list of users
// swagger:response usersResponse
type usersResponseWrapper struct {
	// All current users
	// in: body
	Body []models.User
}

// MODELS USED TO DOCUMENT POST ENDPOINTS:

// Post model
// swagger:response postResponse
type postResponseWrapper struct {
	// Post record
	// in: body
	Body models.Post
}

// List of Posts
// swagger:response postsResponse
type postsResponseWrapper struct {
	// List of post records
	// in: body
	Body []models.Post
}

// Authorization Token
// swagger:response authToken
type tokenResponseWrapper struct {
	// Authorization token
	// in: body
	Body models.Token
}

// ID used for users/posts
// swagger:parameters getUser GetAuthUser GetPost UpdatePost DeletePost
type identifierParamWrapper struct {
	// The identification key
	// in: path
	// required: true
	ID int `json:"id"`
}
