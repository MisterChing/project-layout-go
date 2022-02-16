package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	// welcome
	e.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, this is project-layout-go!")
	})
	//性能分析
	RegisterPprof(e)
	//Metrics
	RegisterMetrics(e)

	//业务router
	RegisterHelloWorld(e)

}
