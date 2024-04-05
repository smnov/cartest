package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Server struct {
	addr   string
	db     Database
	logger *slog.Logger
}

func NewServer(addr string) *Server {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db := MockDB{}
	return &Server{
		addr:   addr,
		db:     db,
		logger: l,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v ...any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func HTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/get/{id}", HTTPHandleFunc(s.GetHandler))
	router.HandleFunc("/delete/{id}", HTTPHandleFunc(s.GetHandler))
	router.HandleFunc("/edit/{id}", HTTPHandleFunc(s.GetHandler))
	router.HandleFunc("/add", HTTPHandleFunc(s.GetHandler))
	s.logger.Info("Server started at port ", s.addr)
	http.ListenAndServe(s.addr, router)
}
