package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nebisin/gowallet/db/model"
)

type Server struct {
	store  *model.Store
	router *gin.Engine
}

func NewServer(store *model.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.transfer)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
