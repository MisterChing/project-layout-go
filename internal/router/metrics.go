package router

import (
	"net/http"
	"project-layout-go/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterMetrics 注册metrics路由
func RegisterMetrics(engine *gin.Engine) {
	router := engine.Group("/").Use(middleware.OnlyLocalIP())
	router.GET("/metrics", gin.WrapF(MetricsHandler))
}

// MetricsHandler metrics路由处理函数
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	handler := promhttp.Handler()
	handler.ServeHTTP(w, r)
}
