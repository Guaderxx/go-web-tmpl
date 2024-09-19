package web

import (
	"github.com/Guaderxx/gowebtmpl/pkg/alog"
	"github.com/Guaderxx/gowebtmpl/pkg/web/middleware"
	"github.com/gin-gonic/gin"
)

type Options struct {
	Port         string
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
	IdleTimeout  int `mapstructure:"idle_timeout"`
	MaxHeaderMB  int `mapstructure:"max_header_mb"`
}

func New(logger alog.ALogger) *gin.Engine {
	router := gin.New()

	router.Use(middleware.Logger(logger))
	router.Use(middleware.Recovery(logger))

	return router
}
