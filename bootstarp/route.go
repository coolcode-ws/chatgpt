package bootstarp

import (
	"sync"

	"github.com//henkgo/chatgpt/routes"
	"github.com/gin-gonic/gin"
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
