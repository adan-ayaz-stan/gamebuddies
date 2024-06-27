package handlers

import (
	"github.com/gin-gonic/gin"
)

func HeartbeatRoute(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
