package web

import (
	"net/http"

	"github.com/Guaderxx/gowebtmpl/pkg/core"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/entity"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/usecase"
	"github.com/Guaderxx/gowebtmpl/pkg/web/routers"
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

func NewTask(r *gin.RouterGroup, core *core.Core) {
	tr := entity.NewTaskRepo(core)
	tc := routers.Task{
		Core:    core,
		Usecase: usecase.NewTaskUsecase(tr),
	}

	r.POST("/task", tc.Create)
}

func NewUser(r *gin.RouterGroup, core *core.Core) {
	ur := entity.NewUserRepo(core)
	uc := routers.User{
		Core:    core,
		Usecase: usecase.NewUserUsecase(ur),
	}

	r.GET("/users", uc.Users)
	r.PUT("/user", uc.UpdateUserName)
	r.DELETE("/user", uc.DeleteByID)

	r.GET("/tasks", uc.UserTasks)
}
