package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (serv *Server) Router(route *chi.Mux) {
	route.Use(middleware.Logger)
	route.Get("/", serv.HelloWorld)
}