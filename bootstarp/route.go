package bootstarp

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/henkgo/chatgpt/routes"
)

var (
	once   sync.Once
	router *gin.Engine
)

// SetUpRoute ...
func SetUpRoute() {
	once.Do(func() {
		router = gin.Default()
		routes.RegisterWebRoutes(router)
	})
}
