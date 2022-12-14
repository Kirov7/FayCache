package http

import (
	cluster "github.com/Kirov7/FayCache/cluster"
	"github.com/Kirov7/FayCache/server/cache"
	"net/http"
)

type Server struct {
	cache.Cache
	cluster.Node
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.Handle("/cluster", s.clusterHandler())
	http.ListenAndServe(s.Addr()+":20131", nil)
}

func NewHTTPServer(c cache.Cache, n cluster.Node) *Server {
	return &Server{c, n}
}
