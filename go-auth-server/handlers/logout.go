package handlers

import (
	"github.com/gin-gonic/gin"
)

func LogoutRoute(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
