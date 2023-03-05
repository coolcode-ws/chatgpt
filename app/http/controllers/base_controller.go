package controllers

import (
	"github.com/gin-gonic/gin"
)

// BaseController ...
type BaseController struct {
}

// ResponseJSON ...
func (*BaseController) ResponseJSON(ctx *gin.Context, code int, errorMsg string, data interface{}) {

	ctx.JSON(code, gin.H{
		"code":     code,
		"errorMsg": errorMsg,
		"data":     data,
	})
	ctx.Abort()
}
