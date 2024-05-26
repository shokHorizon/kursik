package service

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
)

type Courses interface {
	GetCourse(ctx context.Context, id int) (*entity.Course, error)
	GetCourses(ctx context.Context, params entity.Pagination) (entity.Courses, error)
	SearchCourse(ctx context.Context, params entity.SearchParams) (entity.Courses, error)
	PatchCourse(ctx context.Context, course *entity.Course) error
	CreateCourse(ctx context.Context, course *entity.Course) error
	DeleteCourse(ctx context.Context, id int) error
}

type Tasks interface {
	GetTask(ctx context.Context, id int64) (*entity.Task, error)
	GetTasks(ctx context.Context, filter entity.TaskFilter, params entity.Pagination) (entity.Tasks, error)
	SearchTask(ctx context.Context, params entity.SearchParams) (entity.Tasks, error)
	PatchTask(ctx context.Context, task *entity.Task) error
	CreateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int64) error
}
