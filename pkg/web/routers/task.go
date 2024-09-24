package routers

import (
	"net/http"

	"github.com/Guaderxx/gowebtmpl/ent"
	"github.com/Guaderxx/gowebtmpl/pkg/core"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/service"
	"github.com/gin-gonic/gin"
)

type Task struct {
	Core    *core.Core
	Usecase service.Task
}

func (tc *Task) Create(c *gin.Context) {
	var req service.TaskRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}

	res, err := tc.Usecase.Create(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, service.Response[*ent.Task]{
		Code: 200,
		Data: res,
	})
}
