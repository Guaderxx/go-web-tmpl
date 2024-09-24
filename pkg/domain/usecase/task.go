package usecase

import (
	"context"

	"github.com/Guaderxx/gowebtmpl/ent"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/repo"
	"github.com/Guaderxx/gowebtmpl/pkg/domain/service"
)

type task struct {
	taskRepo repo.Task
}

func NewTaskUsecase(tr repo.Task) service.Task {
	return &task{
		taskRepo: tr,
	}
}

func (tu *task) Create(c context.Context, req *service.TaskRequest) (*ent.Task, error) {
	return tu.taskRepo.Create(c, req.Title, req.Status, req.Priority, req.UserID)
}
