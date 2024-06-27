package handlers

import (
	"github.com/gin-gonic/gin"
)

func VerifyEmailRoute(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
