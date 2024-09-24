package repo

import (
	"context"

	"github.com/Guaderxx/gowebtmpl/ent"
)

type Task interface {
	Create(c context.Context, title, status string, priority int, userID uint64) (*ent.Task, error)
	UpdateStatus(c context.Context, status string) (*ent.Task, error)
	UpdatePriority(c context.Context, priority int) (*ent.Task, error)
	DeleteByID(c context.Context, id uint64) error
}
