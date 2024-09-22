package usecase

import (
	"context"

	"github.com/Guaderxx/gowebtmpl/ent"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/repo"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/service"
	"github.com/Guaderxx/gowebtmpl/pkg/web/middleware"
)

type user struct {
	userRepo repo.User
}

func NewUserUsecase(ur repo.User) service.User {
	return &user{
		userRepo: ur,
	}
}

func (uu *user) Create(c context.Context, name, pwd, email string) (*ent.User, error) {
	return uu.userRepo.Create(c, name, pwd, email)
}

func (uu *user) GetUserByID(c context.Context, id uint64) (*ent.User, error) {
	return uu.userRepo.GetByID(c, id)
}

func (uu *user) GetUserByEmail(c context.Context, email string) (*ent.User, error) {
	return uu.userRepo.GetByEmail(c, email)
}

func (uu *user) CreateAccessToken(user *ent.User, secret string, expiry int) (string, error) {
	return middleware.CreateAccessToken(user.ID, user.Name, secret, expiry)
}

func (uu *user) CreateRefreshToken(user *ent.User, secret string, expiry int) (refreshToken string, err error) {
	return middleware.CreateRefreshToken(user.ID, secret, expiry)
}

func (uu *user) ExtractIDFromToken(requestToken, secret string) (uint64, error) {
	return middleware.ExtractIDFromToken(requestToken, secret)
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
