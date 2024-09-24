package service

import (
	"context"

	"github.com/Guaderxx/gowebtmpl/ent"
)

type Task interface {
	Create(c context.Context, req *TaskRequest) (*ent.Task, error)
}

type TaskRequest struct {
	Title    string
	Status   string
	Priority int
	UserID   uint64
}
