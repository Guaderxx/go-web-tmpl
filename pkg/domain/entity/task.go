package entity

import (
	"context"

	"github.com/Guaderxx/gowebtmpl/ent"
	"github.com/Guaderxx/gowebtmpl/ent/task"
	"github.com/Guaderxx/gowebtmpl/pkg/core"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/repo"
)

type taskRepo struct {
	core *core.Core
}

func NewTaskRepo(core *core.Core) repo.Task {
	return &taskRepo{core: core}
}

func (tr *taskRepo) Create(c context.Context, title, status string, priority int, userID uint64) (*ent.Task, error) {
	return tr.core.DB.Task.Create().
		SetTitle(title).
		SetStatus(task.Status(status)).
		SetPriority(priority).
		SetOwnerID(userID).
		Save(c)
}

func (tr *taskRepo) UpdateStatus(c context.Context, status string) (*ent.Task, error) {
	return nil, nil
}

func (tr *taskRepo) UpdatePriority(c context.Context, priority int) (*ent.Task, error) {
	return nil, nil
}

func (tr *taskRepo) DeleteByID(c context.Context, id uint64) error {
	return nil
}
