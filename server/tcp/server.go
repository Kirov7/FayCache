package tcp

import (
	"github.com/Kirov7/FayCache/server/cache"
	"net"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	l, e := net.Listen("tcp", ":12346")
	if e != nil {
		panic(e)
	}
	for {
		c, e := l.Accept()
		if e != nil {
			panic(e)
		}
		go s.process(c)
	}
}

func NewTCPServer(c cache.Cache) *Server {
	return &Server{c}
}
