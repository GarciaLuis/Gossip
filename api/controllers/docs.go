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
// swagger:meta

package controllers

import "github.com/garcialuis/Gossip/api/models"

// A list of users
// swagger:response usersResponse
type usersResponseWrapper struct {
	// All current users
	// in: body
	Body []models.User
}