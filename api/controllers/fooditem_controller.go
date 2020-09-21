package controllers

import (
	"net/http"

	"github.com/garcialuis/Gossip/api/responses"
)

func (server *Server) GetAllFoodItems(w http.ResponseWriter, r *http.Request) {

	foodItems := server.NutriportClient.GetAllFoodItems()

	responses.JSON(w, http.StatusOK, foodItems)
}
