package service

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
)

type Courses interface {
	Get(ctx context.Context, id int64) (*entity.Course, error)
	GetAll(ctx context.Context, params entity.Pagination) (entity.Courses, error)
	Search(ctx context.Context, params entity.SearchParams) (entity.Courses, error)
	Patch(ctx context.Context, course *entity.Course) error
	Create(ctx context.Context, course *entity.Course) error
	Delete(ctx context.Context, id int64) error
}

type Tasks interface {
	Get(ctx context.Context, id int64) (*entity.Task, error)
	GetAll(ctx context.Context, filter entity.TaskFilter, params entity.Pagination) (entity.Tasks, error)
	Search(ctx context.Context, params entity.SearchParams) (entity.Tasks, error)
	Patch(ctx context.Context, task *entity.Task) error
	Create(ctx context.Context, task *entity.Task) error
	Delete(ctx context.Context, id int64) error
}
