package service

import (
	"context"

	"github.com/Guaderxx/gowebtmpl/ent"
)

type User interface {
	Users(c context.Context) (ent.Users, error)
	UpdateNameByID(c context.Context, id uint64, newName string) (*ent.User, error)
	DeleteByID(c context.Context, id uint64) error

	Signup(c context.Context, req *SignupRequest) (*SignupResponse, error)
	Login(c context.Context, req *LoginRequest) (*LoginResponse, error)
	CreateToken(user *ent.User, accessSecret, refreshSecret string, accessExpiry, refreshExpiry int) (accessToken, refreshToken string, err error)
	RefreshToken(c context.Context, req *RefreshTokenRequest) (*RefreshTokenResponse, error)
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
