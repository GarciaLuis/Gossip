package controllers

import (
	"net/http"

	"github.com/garcialuis/Gossip/api/responses"
)

// Home Handler:
// swagger:route GET / home Home
//	Responses:
//		200: description: OK - Welcomes to Gossip API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to Gossip Api")
}
