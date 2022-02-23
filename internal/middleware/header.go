package middleware

import (
    "strings"

    "github.com/gin-gonic/gin"
)

func NoCache() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Header("Cache-Control", "no-store")
    }
}

// Options -
func Options() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        if strings.ToLower(ctx.Request.Method) == "options" {
            ctx.String(200, "ok")
            ctx.Abort()
        }
    }
}
