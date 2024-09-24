package usecase

import (
	"context"
	"errors"

	"github.com/Guaderxx/gowebtmpl/config"
	"github.com/Guaderxx/gowebtmpl/ent"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/repo"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/service"
	"github.com/Guaderxx/gowebtmpl/pkg/web/middleware"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	userRepo repo.User
}

func NewUserUsecase(ur repo.User) service.User {
	return &user{
		userRepo: ur,
	}
}

func (uu *user) Users(c context.Context) (ent.Users, error) {
	return uu.userRepo.Fetch(c)
}

func (uu *user) UpdateNameByID(c context.Context, id uint64, newName string) (*ent.User, error) {
	return uu.userRepo.UpdateName(c, id, newName)
}

func (uu *user) DeleteByID(c context.Context, id uint64) error {
	return uu.userRepo.DeleteByID(c, id)
}

func (uu *user) Signup(c context.Context, req *service.SignupRequest) (*service.SignupResponse, error) {
	_, err := uu.userRepo.GetByEmail(c, req.Email)
	if err == nil {
		return nil, errors.New("user already exists.")
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}
	req.Password = string(encryptedPassword)
	user, err := uu.userRepo.Create(c, req.Name, req.Password, req.Email)
	if err != nil {
		return nil, err
	}

	jwtC, ok := c.Value("c-jwt").(config.JWT)
	if !ok {
		return nil, errors.New("get token failed")
	}

	accessToken, refreshToken, err := uu.CreateToken(
		user,
		jwtC.AccessTokenSecret,
		jwtC.RefreshTokenSecret,
		jwtC.AccessTokenExpiryHour,
		jwtC.RefreshTokenExpiryHour,
	)
	if err != nil {
		return nil, err
	}
	return &service.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uu *user) Login(c context.Context, req *service.LoginRequest) (*service.LoginResponse, error) {
	user, err := uu.userRepo.GetByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}
	// validate password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return nil, errors.New("invalid credentials.")
	}

	jwtC, ok := c.Value("c-jwt").(config.JWT)
	if !ok {
		return nil, errors.New("get token failed")
	}

	accessToken, refreshToken, err := uu.CreateToken(
		user,
		jwtC.AccessTokenSecret,
		jwtC.RefreshTokenSecret,
		jwtC.AccessTokenExpiryHour,
		jwtC.RefreshTokenExpiryHour,
	)
	if err != nil {
		return nil, err
	}

	return &service.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uu *user) CreateToken(user *ent.User, accessSecret, refreshSecret string, accessExpiry, refreshExpiry int) (accessToken, refreshToken string, err error) {
	accessToken, err = middleware.CreateAccessToken(user.ID, user.Name, accessSecret, accessExpiry)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = middleware.CreateRefreshToken(user.ID, refreshSecret, refreshExpiry)
	if err != nil {
		return "", "", err
	}
	return
}

func (uu *user) RefreshToken(c context.Context, req *service.RefreshTokenRequest) (*service.RefreshTokenResponse, error) {
	jwtC, ok := c.Value("c-jwt").(config.JWT)
	if !ok {
		return nil, errors.New("get token failed")
	}

	// Don't get from "x-user-id", validate this token
	id, err := middleware.ExtractIDFromToken(req.RefreshToken, jwtC.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}
	user, err := uu.userRepo.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	accessToken, refreshToken, err := uu.CreateToken(
		user,
		jwtC.AccessTokenSecret,
		jwtC.RefreshTokenSecret,
		jwtC.AccessTokenExpiryHour,
		jwtC.RefreshTokenExpiryHour,
	)
	if err != nil {
		return nil, err
	}
	return &service.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
