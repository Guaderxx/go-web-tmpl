package web

import (
	"net/http"

	"github.com/Guaderxx/gowebtmpl/pkg/core"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/entity"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/usecase"
	"github.com/Guaderxx/gowebtmpl/pkg/web/routers"
	"github.com/gin-gonic/gin"
)

func NewDemo(r *gin.RouterGroup, co *core.Core) {
	r.GET("/ping", Pong(co))
}

func Pong(co *core.Core) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": "pong",
		})
	}
}

func NewUserAccount(r *gin.RouterGroup, core *core.Core) {
	ur := entity.NewUserRepo(core)
	uc := routers.User{
		Core:    core,
		Usecase: usecase.NewUserUsecase(ur),
	}

	r.POST("/signup", uc.Signup)
	r.POST("/login", uc.Login)
	r.POST("/refresh", uc.RefreshToken)

	r.GET("/users", uc.Users)
	r.PUT("/user", uc.UpdateUserName)
	r.DELETE("/user", uc.DeleteByID)
}
