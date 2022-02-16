package router

import (
	"net/http/pprof"
	"project-layout-go/internal/infrastructure/common/config"
	"project-layout-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

const (
	// DefaultPprofPrefix 默认pprof路由前缀
	DefaultPprofPrefix = "/debug/pprof"
)

// RegisterPprof 注册pprof路由
func RegisterPprof(engine *gin.Engine) {
	//pprof
	isOpenpprof := config.AppCfg.GetString("pprof.enable")
	if isOpenpprof == "true" {
		router := engine.Group(DefaultPprofPrefix).Use(middleware.OnlyLocalIP())
		router.GET("/", gin.WrapF(pprof.Index))
		router.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		router.GET("/profile", gin.WrapF(pprof.Profile))
		router.POST("/symbol", gin.WrapF(pprof.Symbol))
		router.GET("/symbol", gin.WrapF(pprof.Symbol))
		router.GET("/trace", gin.WrapF(pprof.Trace))
		router.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		router.GET("/block", gin.WrapH(pprof.Handler("block")))
		router.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		router.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		router.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		router.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}

}
