// Package classification of Gossip API
//
// Documentation for Gossip API
//
//		Schemes: http
//		BasePath: /
//		Version: 1.0.0
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
