package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	server := &Server{}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// enable CORS
	router.Use(CORSMiddleware())

	// Set trusted proxies to localhost
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	// routes
	router.GET("/calculate-packs/:ordersize", server.calculatePacksHandler)

	server.router = router
}

// run HTTP server on specific address to listen for requests
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// CORSMiddleware sets the CORS headers to allow all origins
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
