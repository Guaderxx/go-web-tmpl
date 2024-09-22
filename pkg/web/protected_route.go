package web

import (
	"net/http"

	"github.com/Guaderxx/gowebtmpl/pkg/core"
	"github.com/gin-gonic/gin"
)

func NewProtected(r *gin.RouterGroup, co *core.Core) {
	r.GET("/pong", Pon2(co))
}

func Pon2(co *core.Core) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": "protected pong",
		})
	}
}
