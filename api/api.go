package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.registerRoutes()
	return s
}
func (s *Server) registerRoutes() {
	s.HandleFunc("/hello", s.getList()).Methods("GET")
}

func (s *Server) getList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		listOfItems := []string{"apple", "mango"}

		if err := json.NewEncoder(writer).Encode(listOfItems); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		return

	}
}
