package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
