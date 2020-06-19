package controllers

import (
	"net/http"

	"github.com/garcialuis/Gossip/api/responses"
)

// Home controller welcomes to API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to Gossip Api")
}
