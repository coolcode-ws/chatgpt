package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henkgo/chatgpt/app/http/controllers"
	middleware "github.com/henkgo/chatgpt/app/middlewares"
)

var chatController = controllers.NewChatController()

// RegisterWebRoutes ...
func RegisterWebRoutes(router *gin.Engine) {
	router.Use(middleware.HanderFunc())
	// index
	router.GET("/", chatController.Index)
	// completion
	router.POST("/completion", chatController.Completion)
}
