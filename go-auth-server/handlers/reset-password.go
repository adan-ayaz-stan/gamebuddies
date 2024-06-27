package handlers

import (
	"github.com/gin-gonic/gin"
)

func ResetPasswordRoute(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
