package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)
	return WriteJSON(w, 200, id)
}

func (s *Server) DeleteHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) EditHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) AddHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
