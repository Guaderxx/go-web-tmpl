package web

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/Guaderxx/gowebtmpl/pkg/alog"
	"github.com/gin-gonic/gin"
)

func New(logger alog.ALogger) *gin.Engine {
	router := gin.New()

	router.Use(Logger(logger))
	router.Use(Recovery(logger))

	return router
}

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

// Recovery
func Recovery(logger alog.ALogger) gin.HandlerFunc {
	logger = logger.WithGroup("recovery")

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}
				headersToStr := strings.Join(headers, "\r\n")
				if brokenPipe {
					logger.Warn("broken-pipe", "error", err, "header", headersToStr)
				} else if gin.IsDebugging() {
					logger.Warn("", "error", err, "header", headersToStr)
				} else {
					logger.Warn("", "error", err, "header", headersToStr)
				}

				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) //nolint: errcheck
					c.Abort()
				} else {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}

// SecurityHeaders  Avoid Host Header Injection related attacks (SSRF, Open Redirection)
func SecurityHeaders() gin.HandlerFunc {
	expectedHost := "localhost:8080"

	return func(c *gin.Context) {
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	}
}

type Options struct {
	Port         string
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
	IdleTimeout  int `mapstructure:"idle_timeout"`
	MaxHeaderMB  int `mapstructure:"max_header_mb"`
}
