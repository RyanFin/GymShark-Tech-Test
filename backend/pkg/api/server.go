package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // This is the correct import
	ginSwagger "github.com/swaggo/gin-swagger"

	// DO NOT FORGET TO ADD DOCS PATH!
	// Get the prefix from the go.,mod file
	_ "GymShark-Tech-Test/docs"
)

type Server struct {
	router *gin.Engine
	item   Item
}

func NewServer() (*Server, error) {

	// Create an Items object with the provided name and price
	item := NewItem("gymshark-vest-top", 15.99)
	server := &Server{item: *item}

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
	// Swagger documentation path
	// http://localhost:8080/docs/index.html
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/calculate-packs/:ordersize", server.calculatePacksHandler)
	router.GET("/view-packsizes", server.getPackSizesHandler)
	router.POST("/add-packsize", server.addPackSizeHandler)
	router.DELETE("/remove-packsize", server.removePackSizeHandler)

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
