package Tasks

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
	"github.com/shokHorizon/kursik/internal/service"
)

type UseCase struct {
	svc service.Tasks
}

func NewUseCase(svc service.Tasks) *UseCase {
	return &UseCase{svc: svc}
}

func (uc *UseCase) CreateTask(ctx context.Context, task *entity.Task) error {
	return uc.svc.CreateTask(ctx, task)
}

func (uc *UseCase) GetTasks(ctx context.Context, filter entity.TaskFilter, pagination entity.Pagination) (entity.Tasks, error) {
	return uc.svc.GetTasks(ctx, filter, pagination)
}

func (uc *UseCase) GetTask(ctx context.Context, id int64) (*entity.Task, error) {
	return uc.svc.GetTask(ctx, id)
}

func (uc *UseCase) PatchTask(ctx context.Context, task *entity.Task) error {
	return uc.svc.PatchTask(ctx, task)
}

func (uc *UseCase) DeleteTask(ctx context.Context, id int64) error {
	return uc.svc.DeleteTask(ctx, id)
}

func (uc *UseCase) SearchTask(ctx context.Context, params entity.SearchParams) (entity.Tasks, error) {
	return uc.svc.SearchTask(ctx, params)
}
