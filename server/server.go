package server

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	server *http.Server
}

// NewServer creates a new server.
func NewServer(port string) Server {
	ginHandler := gin.Default() // TODO: change to gin.New()

	srv := server{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: ginHandler,
		},
	}

	// TODO ginHandler should look something like this:
	// // Setup logging, metrics, cors middlewares.
	// r.Use(
	// 	// Ignore logging healthcheck routes.
	// 	gin.LoggerWithWriter(gin.DefaultWriter, healthcheckRoute),
	// 	gin.Recovery(),
	// 	apmgin.Middleware(r),
	// 	cors.New(corsRouterConfig()),
	// 	// Elasticsearch logger middleware.
	// 	loggermiddleware.SetLogger(
	// 		&loggermiddleware.Config{
	// 			Logger:             logger,
	// 			SkipPath:           []string{healthcheckRoute},
	// 			SkipBodyPathRegexp: regexp.MustCompile(uploadRouteRegexp),
	// 		},
	// 	),
	// )

	ginHandler.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "DELETE", "POST"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           0,
	}))

	ginHandler.Use(ErrorHandler)

	UseServerRouter(ginHandler)

	return &srv
}

// Run starts the server.
func (s *server) Run() {
	log.Printf("Server started on port %v\n", s.server.Addr)

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// Serve creates a server and starts it.
func Serve(port string) {
	server := NewServer(port)
	server.Run()
}
