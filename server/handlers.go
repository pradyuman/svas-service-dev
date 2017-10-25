package server

import (
	"github.com/gin-gonic/gin"
)

type Handlers struct{}

func (h *Handlers) GetStatus(ctx *gin.Context) {
	Success(ctx, gin.H{
		"status": 200,
	})
	// Get Status
}
