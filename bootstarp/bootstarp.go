package bootstarp

import (
	"net/http"

	"github.com/henkgo/chatgpt/common/config"
	"github.com/henkgo/chatgpt/common/logger"
)

// StartWebChatServer :
func StartWebChatServer() {
	SetUpRoute()
	initTemplateDir()
	initStaticServer()

	if err := router.Run(":" + config.LoadConfig().Addr); err != nil {
		logger.Error("run webserver error %s", err)
		return
	}
}

// initTemplate : html resources path
func initTemplateDir() {
	router.LoadHTMLGlob("resources/view/*")
}

// initStaticServer : static resource resources
func initStaticServer() {
	router.StaticFS("resources/static", http.Dir("static"))
	router.StaticFile("logo192.png", "static/logo192.png")
	router.StaticFile("logo512.png", "static/logo512.png")
}
