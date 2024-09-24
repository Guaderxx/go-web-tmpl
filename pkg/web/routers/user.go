package routers

import (
	"context"
	"net/http"

	"github.com/Guaderxx/gowebtmpl/ent"
	"github.com/Guaderxx/gowebtmpl/pkg/core"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/service"
	"github.com/gin-gonic/gin"
)

type User struct {
	Core    *core.Core
	Usecase service.User
}

func (uc *User) Signup(c *gin.Context) {
	var req service.SignupRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  400,
			Error: err.Error(),
		})
		return
	}

	ctx := context.WithValue(c, "c-jwt", uc.Core.Config.JWT)

	res, err := uc.Usecase.Signup(ctx, &req)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, service.Response[service.SignupResponse]{
		Code: 200,
		Data: *res,
	})
}

func (uc *User) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}

	ctx := context.WithValue(c, "c-jwt", uc.Core.Config.JWT)

	res, err := uc.Usecase.Login(ctx, &req)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, service.Response[service.LoginResponse]{
		Code: 200,
		Data: *res,
	})
}

func (uc *User) RefreshToken(c *gin.Context) {
	var req service.RefreshTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  400,
			Error: err.Error(),
		})
		return
	}

	ctx := context.WithValue(c, "c-jwt", uc.Core.Config.JWT)

	res, err := uc.Usecase.RefreshToken(ctx, &req)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusOK, service.Response[service.RefreshTokenResponse]{
		Code: 200,
		Data: *res,
	})
}

func (uc *User) Users(c *gin.Context) {
	users, err := uc.Usecase.Users(c)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, service.Response[ent.Users]{
		Code: 200,
		Data: users,
	})
}

func (uc *User) UpdateUserName(c *gin.Context) {
	type Req struct {
		ID   uint64
		Name string
	}
	var req Req
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  400,
			Error: err.Error(),
		})
		return
	}
	user, err := uc.Usecase.UpdateNameByID(c, req.ID, req.Name)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  400,
			Error: err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := uc.Usecase.CreateToken(
		user,
		uc.Core.Config.JWT.AccessTokenSecret,
		uc.Core.Config.JWT.RefreshTokenSecret,
		uc.Core.Config.JWT.AccessTokenExpiryHour,
		uc.Core.Config.JWT.RefreshTokenExpiryHour,
	)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}
	type Resp struct {
		*ent.User
		service.LoginResponse
	}

	c.JSON(http.StatusOK, service.Response[Resp]{
		Code: 200,
		Data: Resp{
			User: user,
			LoginResponse: service.LoginResponse{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		},
	})
}

func (uc *User) DeleteByID(c *gin.Context) {
	type Req struct {
		ID uint64
	}
	var req Req
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  400,
			Error: err.Error(),
		})
		return
	}
	err := uc.Usecase.DeleteByID(c, req.ID)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  400,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, service.Response[string]{
		Code: 200,
		Data: "delete user succeed",
	})
}

func (uc *User) UserTasks(c *gin.Context) {
	uid := c.GetUint64("x-user-id")
	tasks, err := uc.Usecase.UserTasks(c, uid)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, service.Response[ent.Tasks]{
		Code: 200,
		Data: tasks,
	})
}
