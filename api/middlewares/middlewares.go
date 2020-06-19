package middlewares

import (
	"errors"
	"net/http"

	"github.com/garcialuis/Gossip/api/auth"
	"github.com/garcialuis/Gossip/api/responses"
)

// SetMiddlewareJSON is used to set the content type as JSON
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication makes use of auth to check if token is valid
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
