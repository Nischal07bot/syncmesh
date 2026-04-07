package server

import (
	"log"
	"net/http"
	"github.com/Nischal07bot/syncmesh/internal/router"
)

type Server struct {
	Addr   string
}

func New(addr string) *Server {
	return &Server{
		Addr:   addr,
	}
}

func (s *Server) Start() error {
	mux:= router.SetupRoutes()
	log.Printf("Starting server on %s\n", s.Addr)
	return http.ListenAndServe(s.Addr, mux)
}