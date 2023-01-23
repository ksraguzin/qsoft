package handlers

import (
	"github.com/gin-gonic/gin"
)

type handler struct {
}

func PingMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		if c.GetHeader("X-PING") == "ping" {
			c.Writer.Header().Set("X-PONG", "pong")
			c.Next()
		}
	}
}
func RegisterRoutes(r *gin.Engine) {
	h := &handler{}
	r.Use(PingMiddleware())
	r.GET("/when/:year", h.GetDays)
}
