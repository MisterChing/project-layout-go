package router

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterHelloWorld(e *gin.Engine) {
    router := e.Group("/")
    router.GET("/hello", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })
}
