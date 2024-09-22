package entity

import (
	"context"
	"errors"

	"github.com/Guaderxx/gowebtmpl/ent"
	"github.com/Guaderxx/gowebtmpl/ent/user"
	"github.com/Guaderxx/gowebtmpl/pkg/core"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/repo"
)

type userRepo struct {
	core *core.Core
}

func NewUserRepo(core *core.Core) repo.User {
	return &userRepo{core: core}
}

func (ur *userRepo) Create(ctx context.Context, name, pwd, email string) (*ent.User, error) {
	u, err := ur.core.DB.User.Create().
		SetName(name).
		SetPassword(pwd).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		ur.core.Logger.Warn("user created failed", "error", err.Error())
		return nil, errors.New("user created failed")
	}
	ur.core.Logger.Info("user created succeed", "user", u)
	return u, nil
}

func (ur *userRepo) Fetch(ctx context.Context) (ent.Users, error) {
	users, err := ur.core.DB.User.Query().
		All(ctx)
	if err != nil {
		ur.core.Logger.Warn("fetch users failed", "error", err.Error())
		return nil, errors.New("fetch users failed")
	}
	return users, nil
}

func (ur *userRepo) GetByID(ctx context.Context, id uint64) (*ent.User, error) {
	user, err := ur.core.DB.User.Get(ctx, id)
	if err != nil {
		ur.core.Logger.Warn("get user by id failed", "error", err.Error(), "id", id)
		return nil, errors.New("get user failed")
	}
	return user, nil
}

func (ur *userRepo) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	user, err := ur.core.DB.User.Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		ur.core.Logger.Warn("get user by email failed", "error", err.Error(), "email", email)
		return nil, errors.New("get user failed")
	}

	return user, nil
}

func (ur *userRepo) UpdateName(ctx context.Context, id uint64, newName string) (*ent.User, error) {
	return ur.core.DB.User.UpdateOneID(id).
		SetName(newName).
		Save(ctx)
}

// DeleteByID  Cause the soft-delete, this will only update the `deleted_at` column, not actually *delete* the row.
func (ur *userRepo) DeleteByID(ctx context.Context, id uint64) error {
	return ur.core.DB.User.
		DeleteOneID(id).
		Exec(ctx)
}
