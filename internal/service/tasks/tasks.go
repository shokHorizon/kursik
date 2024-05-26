package tasks

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
	"github.com/shokHorizon/kursik/internal/service"
)

var _ service.Tasks = (*Service)(nil)

type Repository interface {
	GetTask(ctx context.Context, id int64) (*entity.Task, error)
	GetTasks(ctx context.Context, filter entity.TaskFilter, params entity.Pagination) (entity.Tasks, error)
	SearchTask(ctx context.Context, params entity.SearchParams) (entity.Tasks, error)
	PatchTask(ctx context.Context, task *entity.Task) error
	CreateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int64) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetTask(ctx context.Context, id int64) (*entity.Task, error) {
	return s.repository.GetTask(ctx, id)
}

func (s *Service) GetTasks(ctx context.Context, filter entity.TaskFilter, params entity.Pagination) (entity.Tasks, error) {
	return s.repository.GetTasks(ctx, filter, params)
}

func (s *Service) SearchTask(ctx context.Context, params entity.SearchParams) (entity.Tasks, error) {
	return s.repository.SearchTask(ctx, params)
}

func (s *Service) PatchTask(ctx context.Context, task *entity.Task) error {
	return s.repository.PatchTask(ctx, task)
}

func (s *Service) CreateTask(ctx context.Context, task *entity.Task) error {
	return s.repository.CreateTask(ctx, task)
}

func (s *Service) DeleteTask(ctx context.Context, id int64) error {
	return s.repository.DeleteTask(ctx, id)
}
