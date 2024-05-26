package courses

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
	service "github.com/shokHorizon/kursik/internal/serivce"
)

// Check Service on service.course implementation
var _ service.Courses = (*Service)(nil)

type Repository interface {
	Get(ctx context.Context, id int64) (*entity.Course, error)
	GetAll(ctx context.Context, params entity.Pagination) (entity.Courses, error)
	Search(ctx context.Context, params entity.SearchParams) (entity.Courses, error)
	Patch(ctx context.Context, course *entity.Course) error
	Create(ctx context.Context, course *entity.Course) error
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

func (s *Service) Get(ctx context.Context, id int64) (*entity.Course, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) GetAll(ctx context.Context, params entity.Pagination) (entity.Courses, error) {
	return s.repository.GetAll(ctx, params)
}

func (s *Service) Search(ctx context.Context, params entity.SearchParams) (entity.Courses, error) {
	return s.repository.Search(ctx, params)
}

func (s *Service) Patch(ctx context.Context, course *entity.Course) error {
	return s.repository.Patch(ctx, course)
}

func (s *Service) Create(ctx context.Context, course *entity.Course) error {
	return s.repository.Create(ctx, course)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}
