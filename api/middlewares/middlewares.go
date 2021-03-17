package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jiprakoso/latihan_go/api/auth"
	"github.com/jiprakoso/latihan_go/api/responses"
)

//SetmiddlewareJSON public func, This will format all responses to JSON
func SetmiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//SetMiddlewareAuthentication public func, This will check for the validity of the authentication token provided.
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			fmt.Println("SetMiddlewareAuthentication")
			return
		}
		next(w, r)
	}
}
