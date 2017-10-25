package server

import (
	"github.com/gin-gonic/gin"
)

type Handlers struct{
	Status Status
}

func (h *Handlers) GetStatus(ctx *gin.Context) {
	Success(ctx, h.Status)
}
