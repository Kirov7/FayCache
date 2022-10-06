package http

import (
	"github.com/Kirov7/FayCache/server/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.ListenAndServe(":20131", nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
