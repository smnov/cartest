package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/smnov/cartest/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Server struct {
	addr   string
	db     Database
	logger *slog.Logger
}

func NewServer(addr string, db Database, logger *slog.Logger) *Server {
	return &Server{
		addr:   addr,
		db:     db,
		logger: logger,
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
	// init swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	// init routes
	router.HandleFunc("/cars/get", HTTPHandleFunc(s.GetCarsHandler)).Methods("GET")
	router.HandleFunc("/cars/delete/{id}", HTTPHandleFunc(s.DeleteCarHandler)).Methods("DELETE")
	router.HandleFunc("/cars/update/{id}", HTTPHandleFunc(s.UpdateCarHandler)).Methods("PATCH")
	router.HandleFunc("/cars/add", HTTPHandleFunc(s.AddCarHandler)).Methods("POST")
	s.logger.Info("Starting server...", "port", s.addr)
	err := http.ListenAndServe(s.addr, handlers.CORS()(router))
	if err != nil {
		s.logger.Error("Server failed to start", "error", err)
	}
}
