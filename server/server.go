package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)


type Server struct {
	lg *zap.Logger
	Port string
}

func NewServer(lg *zap.Logger) *Server {
	return &Server{
		lg: lg,
		Port: ":9999",
	}
}

func (serv *Server) StartServer() {
	route := chi.NewRouter()
	serv.Router(route)

	if err := http.ListenAndServe(serv.Port, route); err != nil {
		serv.lg.Fatal("Server can't start",
		zap.Error(err),
		)
	}
	
}