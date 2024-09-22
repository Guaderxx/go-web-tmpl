package repo

import (
	"context"

	"github.com/Guaderxx/gowebtmpl/ent"
)

type User interface {
	Create(c context.Context, name, pwd, email string) (*ent.User, error)
	Fetch(c context.Context) (ent.Users, error)
	GetByID(c context.Context, id uint64) (*ent.User, error)
	GetByEmail(c context.Context, email string) (*ent.User, error)

	UpdateName(c context.Context, id uint64, name string) (*ent.User, error)
	DeleteByID(c context.Context, id uint64) error
}
