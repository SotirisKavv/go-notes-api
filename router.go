package main

import (
	"net/http"
	"notes-api/middleware"

	"github.com/gorilla/mux"
)

const noteByIDRoute = "/note/{id}"

func NewRouter(h http.Handler) *mux.Router {
	server := mux.NewRouter()
	server.Use(middleware.Auth)
	server.Use(middleware.Log)

	server.Handle("/notes", h).Methods("GET")
	server.Handle("/note", h).Methods("POST")
	server.Handle(noteByIDRoute, h).Methods("GET")
	server.Handle(noteByIDRoute, h).Methods("PUT")
	server.Handle(noteByIDRoute, h).Methods("DELETE")

	return server
}
