package routers

import (
	"net/http"

	"github.com/Guaderxx/gowebtmpl/ent"
	"github.com/Guaderxx/gowebtmpl/pkg/core"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	_, err := uc.Usecase.GetUserByEmail(c, req.Email)
	if err == nil {
		c.JSON(http.StatusOK, service.Response[string]{
			Code: 200,
			Data: "user already exists.",
		})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}

	req.Password = string(encryptedPassword)
	user, err := uc.Usecase.Create(c, req.Name, req.Password, req.Email)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}

	accessToken, err := uc.Usecase.CreateAccessToken(
		user,
		uc.Core.Config.JWT.AccessTokenSecret,
		uc.Core.Config.JWT.AccessTokenExpiryHour,
	)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}
	refreshToken, err := uc.Usecase.CreateRefreshToken(
		user,
		uc.Core.Config.JWT.RefreshTokenSecret,
		uc.Core.Config.JWT.RefreshTokenExpiryHour,
	)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, service.Response[service.SignupResponse]{
		Code: 200,
		Data: service.SignupResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
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
	user, err := uc.Usecase.GetUserByEmail(c, req.Email)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}
	// validate password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  401,
			Error: "invalid credentials.",
		})
		return
	}

	accessToken, err := uc.Usecase.CreateAccessToken(
		user,
		uc.Core.Config.JWT.AccessTokenSecret,
		uc.Core.Config.JWT.AccessTokenExpiryHour,
	)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}
	refreshToken, err := uc.Usecase.CreateRefreshToken(
		user,
		uc.Core.Config.JWT.RefreshTokenSecret,
		uc.Core.Config.JWT.RefreshTokenExpiryHour,
	)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, service.Response[service.LoginResponse]{
		Code: 200,
		Data: service.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
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

	id, err := uc.Usecase.ExtractIDFromToken(req.RefreshToken, uc.Core.Config.JWT.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  401,
			Error: "user not found",
		})
		return
	}
	user, err := uc.Usecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  401,
			Error: "user not found",
		})
		return
	}
	accessToken, err := uc.Usecase.CreateAccessToken(
		user,
		uc.Core.Config.JWT.AccessTokenSecret,
		uc.Core.Config.JWT.AccessTokenExpiryHour,
	)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}
	refreshToken, err := uc.Usecase.CreateRefreshToken(
		user,
		uc.Core.Config.JWT.RefreshTokenSecret,
		uc.Core.Config.JWT.RefreshTokenExpiryHour,
	)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, service.Response[service.RefreshTokenResponse]{
		Code: 200,
		Data: service.RefreshTokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}

func (uc *User) Users(c *gin.Context) {
	users, err := uc.Usecase.Users(c)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: "get users failed",
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

	accessToken, err := uc.Usecase.CreateAccessToken(
		user,
		uc.Core.Config.JWT.AccessTokenSecret,
		uc.Core.Config.JWT.AccessTokenExpiryHour,
	)
	if err != nil {
		c.JSON(http.StatusOK, service.ErrorResponse{
			Code:  500,
			Error: err.Error(),
		})
		return
	}
	refreshToken, err := uc.Usecase.CreateRefreshToken(
		user,
		uc.Core.Config.JWT.RefreshTokenSecret,
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
