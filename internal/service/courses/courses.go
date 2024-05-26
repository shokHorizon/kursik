package courses

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
	"github.com/shokHorizon/kursik/internal/service"
)

// Check Service on service.course implementation
var _ service.Courses = (*Service)(nil)

type Repository interface {
	GetCourse(ctx context.Context, id int) (*entity.Course, error)
	GetCourses(ctx context.Context, params entity.Pagination) (entity.Courses, error)
	SearchCourse(ctx context.Context, params entity.SearchParams) (entity.Courses, error)
	PatchCourse(ctx context.Context, course *entity.Course) error
	CreateCourse(ctx context.Context, course *entity.Course) error
	DeleteCourse(ctx context.Context, id int) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetCourse(ctx context.Context, id int) (*entity.Course, error) {
	return s.repository.GetCourse(ctx, id)
}

func (s *Service) GetCourses(ctx context.Context, params entity.Pagination) (entity.Courses, error) {
	return s.repository.GetCourses(ctx, params)
}

func (s *Service) SearchCourse(ctx context.Context, params entity.SearchParams) (entity.Courses, error) {
	return s.repository.SearchCourse(ctx, params)
}

func (s *Service) PatchCourse(ctx context.Context, course *entity.Course) error {
	return s.repository.PatchCourse(ctx, course)
}

func (s *Service) CreateCourse(ctx context.Context, course *entity.Course) error {
	return s.repository.CreateCourse(ctx, course)
}

func (s *Service) DeleteCourse(ctx context.Context, id int) error {
	return s.repository.DeleteCourse(ctx, id)
}
