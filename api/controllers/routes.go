package controllers

import "github.com/jiprakoso/latihan_go/api/middlewares"

func (s *Server) initializeRoutes() {
	// Home route
	s.Router.HandleFunc("/", middlewares.SetmiddlewareJSON(s.Home)).Methods("GET")

	//login Route
	s.Router.HandleFunc("/login", middlewares.SetmiddlewareJSON(s.Login)).Methods("POST")

	//User Route
	s.Router.HandleFunc("/users", middlewares.SetmiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetmiddlewareJSON(s.GetAllUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetmiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetmiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetmiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetmiddlewareJSON(s.GetAllPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetmiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetmiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
