package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(h *Handlers) {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	server.GET("/status", h.GetStatus)
	//server.POST("/status", h.PostStatus)
	//server.GET("/request", h.GetRequest)
	//server.POST("/request", h.PostRequest)

	log.Printf("Server listening on port 8080")
	server.Run(":8080")
}
