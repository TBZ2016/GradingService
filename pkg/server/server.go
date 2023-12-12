package server

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port int
	// Add any other configuration options as needed
}

type Server struct {
	Router *gin.Engine
	config Config
}

func NewServer(config Config) *Server {
	return &Server{
		Router: gin.Default(),
		config: config,
	}
}

func (s *Server) Start() error {
	return s.Router.Run(":" + strconv.Itoa(s.config.Port))
}

func (s *Server) AddMiddleware(middleware gin.HandlerFunc) {
	s.Router.Use(middleware)
}
