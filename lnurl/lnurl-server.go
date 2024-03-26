package lnurl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterRunOnServer() {
	router := setupRouterOnServer()
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func setupRouterOnServer() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}
