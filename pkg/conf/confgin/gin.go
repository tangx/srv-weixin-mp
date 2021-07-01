package confgin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Address string `env:""`
	Port    int    `env:""`

	e *gin.Engine
}

func (s *Server) SetDefaults() {

	if s.Port == 0 {
		s.Port = 80
	}
}
func (s *Server) Initial() {
	s.SetDefaults()

	if s.e == nil {
		s.e = gin.Default()
	}
}
func (s *Server) WithBaseRouter(fn func(e *gin.Engine)) *Server {
	fn(s.e)
	return s
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", s.Address, s.Port)
	return s.e.Run(addr)
}
