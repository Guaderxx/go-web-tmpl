package middleware

import (
	"time"

	"github.com/Guaderxx/gowebtmpl/pkg/alog"
	"github.com/gin-gonic/gin"
)

// Logger  Replace the default gin logger.
func Logger(logger alog.ALogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// process request
		c.Next()

		// after request
		latency := time.Since(start)
		if raw != "" {
			path = path + "?" + raw
		}
		logger.Info("request", "path", path,
			"method", c.Request.Method,
			"status-code", c.Writer.Status(),
			"body-size", c.Writer.Size(),
			"client-ip", c.ClientIP(),
			"latency", latency,
			"error-message", c.Errors.ByType(gin.ErrorTypePrivate).String(),
			// "keys", c.Keys,
		)
	}
}
