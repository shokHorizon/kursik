package Courses

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
	"github.com/shokHorizon/kursik/internal/service"
)

type UseCase struct {
	svc service.Courses
}

func NewUseCase(svc service.Courses) *UseCase {
	return &UseCase{svc: svc}
}

func (uc *UseCase) CreateCourse(ctx context.Context, course *entity.Course) error {
	return uc.svc.CreateCourse(ctx, course)
}

func (uc *UseCase) GetCourses(ctx context.Context, pagination entity.Pagination) (entity.Courses, error) {
	return uc.svc.GetCourses(ctx, pagination)
}

func (uc *UseCase) GetCourse(ctx context.Context, id int) (*entity.Course, error) {
	return uc.svc.GetCourse(ctx, id)
}

func (uc *UseCase) PatchCourse(ctx context.Context, course *entity.Course) error {
	return uc.svc.PatchCourse(ctx, course)
}

func (uc *UseCase) DeleteCourse(ctx context.Context, id int) error {
	return uc.svc.DeleteCourse(ctx, id)
}

func (uc *UseCase) SearchCourse(ctx context.Context, params entity.SearchParams) (entity.Courses, error) {
	return uc.svc.SearchCourse(ctx, params)
}
