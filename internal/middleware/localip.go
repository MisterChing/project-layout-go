package middleware

import (
    "net/http"
    "project-layout-go/pkg/netx"

    "github.com/gin-gonic/gin"
)

// OnlyLocalIP 仅限制内网访问
func OnlyLocalIP() gin.HandlerFunc {
    return func(c *gin.Context) {
        clientIP := c.ClientIP()
        //ipv6 回环 ::1
        if clientIP == "::1" {
            clientIP = "127.0.0.1"
        }
        if netx.IsLan(clientIP) {
            c.Next()
        } else {
            c.AbortWithStatus(http.StatusForbidden)
            return
        }
    }
}
