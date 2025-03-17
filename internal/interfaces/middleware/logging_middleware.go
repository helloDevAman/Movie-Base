package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs each request with method, URL, and response time
func LoggingMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Log request start
		log.Printf("[Gin] Started %s %s", c.Request.Method, c.Request.URL.Path)

		// Process request
		next(c)

		// Log request end
		log.Printf("[Gin] Completed %s %s | Duration: %s", c.Request.Method, c.Request.URL.Path, time.Since(start))
	}
}
