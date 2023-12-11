package server

import (
	"github.com/gin-gonic/gin"
)

// Config holds the configuration options for the server.
type Config struct {
	Port int
	// Add any other configuration options as needed
}

// Server represents an HTTP server.
type Server struct {
	Router *gin.Engine
	config Config
}

// NewServer creates a new instance of the Server.
func NewServer(config Config) *Server {
	return &Server{
		Router: gin.Default(), // Use Gin's default router
		config: config,
	}
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	return s.Router.Run(":" + string(rune(s.config.Port)))
}

// AddMiddleware adds middleware to the server's router.
func (s *Server) AddMiddleware(middleware gin.HandlerFunc) {
	s.Router.Use(middleware)
}
