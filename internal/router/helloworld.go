package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHelloWorld(e *gin.Engine) {
	router := e.Group("/")
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
}
