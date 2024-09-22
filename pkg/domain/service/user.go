package service

import (
	"context"

	"github.com/Guaderxx/gowebtmpl/ent"
)

type User interface {
	Create(c context.Context, name, pwd, email string) (*ent.User, error)
	GetUserByID(c context.Context, id uint64) (*ent.User, error)
	GetUserByEmail(c context.Context, email string) (*ent.User, error)
	CreateAccessToken(user *ent.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *ent.User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (uint64, error)

	Users(c context.Context) (ent.Users, error)
	UpdateNameByID(c context.Context, id uint64, newName string) (*ent.User, error)
	DeleteByID(c context.Context, id uint64) error
}

type SignupRequest struct {
	Name     string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required,password"`
}

type SignupResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Email    string `binding:"required"`
	Password string `binding:"required,password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `binding:"required,jwt"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
