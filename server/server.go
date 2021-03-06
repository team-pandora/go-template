package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/team-pandora/go-template/utils"
)

// Server is the interface for the server.
type Server interface {
	Run()
}

type server struct {
	server *http.Server
}

// NewServer creates a new server.
func NewServer(port string) Server {
	gin.DisableConsoleColor()
	ginHandler := gin.New()

	srv := server{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: ginHandler,
		},
	}

	ginHandler.Use(
		LoggerMiddleware(),
		errorHandler,
		gin.Recovery(),
		cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "PUT", "DELETE", "POST"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowHeaders:     []string{"Content-Type"},
			AllowCredentials: true,
			MaxAge:           0,
		}),
	)

	useServerRouter(ginHandler)

	return &srv
}

// Run starts the server.
func (s *server) Run() {
	utils.Log.Infof("Server started on port %v\n", s.server.Addr)

	if err := s.server.ListenAndServe(); err != nil {
		panic(fmt.Errorf("failed starting server: %v", err))
	}
}

// Serve creates a server and starts it.
func Serve(port string) {
	server := NewServer(port)
	server.Run()
}
