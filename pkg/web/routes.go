package web

import (
	"github.com/Guaderxx/gowebtmpl/pkg/core"
	"github.com/Guaderxx/gowebtmpl/pkg/web/middleware"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine, co *core.Core) {
	publicRouter := r.Group("")
	NewDemo(publicRouter, co)
	NewUserAccount(publicRouter, co)

	protectedRouter := r.Group("")
	protectedRouter.Use(middleware.JwtAuth(co.Config.JWT.AccessTokenSecret))
	NewProtected(protectedRouter, co)
}
