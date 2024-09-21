package web

import (
	"github.com/Guaderxx/gowebtmpl/pkg/alog"
	"github.com/Guaderxx/gowebtmpl/pkg/web/middleware"
	"github.com/gin-gonic/gin"
)

func New(logger alog.ALogger) *gin.Engine {
	router := gin.New()

	router.Use(middleware.Logger(logger))
	router.Use(middleware.Recovery(logger))

	return router
}
