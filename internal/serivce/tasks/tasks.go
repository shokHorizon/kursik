package tasks

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
	service "github.com/shokHorizon/kursik/internal/serivce"
)

var _ service.Tasks = (*Service)(nil)

type Repository interface {
	Get(ctx context.Context, id int64) (entity.Task, error)
	GetAll(ctx context.Context, filter entity.TaskFilter, params entity.Pagination) (entity.Tasks, error)
	Search(ctx context.Context, params entity.SearchParams) (entity.Tasks, error)
	Patch(ctx context.Context, task entity.Task) error
	Create(ctx context.Context, task entity.Task) error
	Delete(ctx context.Context, id int64) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Get(ctx context.Context, id int64) (entity.Task, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) GetAll(ctx context.Context, filter entity.TaskFilter, params entity.Pagination) (entity.Tasks, error) {
	return s.repository.GetAll(ctx, filter, params)
}

func (s *Service) Search(ctx context.Context, params entity.SearchParams) (entity.Tasks, error) {
	return s.repository.Search(ctx, params)
}

func (s *Service) Patch(ctx context.Context, task entity.Task) error {
	return s.repository.Patch(ctx, task)
}

func (s *Service) Create(ctx context.Context, task entity.Task) error {
	return s.repository.Create(ctx, task)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}
