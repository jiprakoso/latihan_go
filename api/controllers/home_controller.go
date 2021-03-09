package controllers

import (
	"net/http"

	"github.com/jiprakoso/latihan_go/api/responses"
)

//Home public method
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To this awesome API")
}
