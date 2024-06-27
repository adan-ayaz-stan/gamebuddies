package handlers

import (
	"github.com/gin-gonic/gin"
)

func LoginRoute(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
