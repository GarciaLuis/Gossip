package controllers

import (
	"net/http"

	"github.com/garcialuis/Gossip/api/middlewares"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) InitializeRoutes() {

	// Home:
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login:
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// Posts:
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	// Users:
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	s.Router.HandleFunc("/users/attributes/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUserAttributes))).Methods("PUT")

	// User - BMI:
	// TODO: Authenticate and Modify to be: /user/bmi/{id} to get the user's bmi with their personalized information (weight, heigth)
	s.Router.HandleFunc("/user/bmi", middlewares.SetMiddlewareJSON(s.GetUserBMI)).Methods("GET")

	// User - TEE:
	s.Router.HandleFunc("/user/tee", middlewares.SetMiddlewareJSON(s.GetUserTEE)).Methods("GET")

	// FoodItems:
	s.Router.HandleFunc("/fooditems", middlewares.SetMiddlewareJSON(s.GetAllFoodItems)).Methods("GET")

	s.Router.HandleFunc("/private/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUsers))).Methods("GET")
	s.Router.HandleFunc("/private/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.AuthenticatedGetUser))).Methods("GET")

	// Swagger Docs:
	opts := middleware.RedocOpts{SpecURL: "../../swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	s.Router.Handle("/docs", sh)
	s.Router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
}
