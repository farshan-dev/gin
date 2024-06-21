package api

import (
	db "github.com/farshan-dev/gin/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for banking services.
type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store * db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.listAccount)

	server.router = router
	return server
}

func (server * Server) Start(address string) error {
	return server.router.Run(address)
}

// return json error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}