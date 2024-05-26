package usecase

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
	Courses2 "github.com/shokHorizon/kursik/internal/usecase/Courses"
)

var (
	_ Courses = (*Courses2.UseCase)(nil)
)

type Courses interface {
	CreateCourse(ctx context.Context, course *entity.Course) error
	GetCourses(ctx context.Context, pagination entity.Pagination) (entity.Courses, error)
	GetCourse(ctx context.Context, id int) (*entity.Course, error)
	PatchCourse(ctx context.Context, course *entity.Course) error
	DeleteCourse(ctx context.Context, id int) error
}

type Tasks interface {
	CreateTask(ctx context.Context, task *entity.Task) error
	GetTasks(ctx context.Context, filter entity.TaskFilter, pagination entity.Pagination) (entity.Tasks, error)
	GetTask(ctx context.Context, id int) (*entity.Task, error)
	PatchTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}
