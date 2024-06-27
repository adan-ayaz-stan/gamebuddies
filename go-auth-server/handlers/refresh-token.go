package handlers

import (
	"github.com/gin-gonic/gin"
)

func RefreshRoute(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
