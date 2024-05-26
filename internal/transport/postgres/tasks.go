package postgres

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
	service "github.com/shokHorizon/kursik/internal/serivce"
	"gorm.io/gorm"
)

var _ service.Tasks = (*TasksRepository)(nil)

type TasksRepository struct {
	db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) *TasksRepository {
	return &TasksRepository{db: db}
}

func (r TasksRepository) Get(ctx context.Context, id int64) (*entity.Task, error) {
	var task entity.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r TasksRepository) GetAll(ctx context.Context, filter entity.TaskFilter, params entity.Pagination) (entity.Tasks, error) {
	var tasks entity.Tasks
	err := r.db.WithContext(ctx).Limit(int(*params.Limit)).Offset(int(*params.Offset)).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r TasksRepository) Search(ctx context.Context, params entity.SearchParams) (entity.Tasks, error) {
	var tasks entity.Tasks
	err := r.db.WithContext(ctx).
		Where("name LIKE?", "%"+*params.Search+"%").Limit(int(*params.Limit)).Offset(int(*params.Offset)).
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r TasksRepository) Patch(ctx context.Context, task *entity.Task) error {
	return r.db.WithContext(ctx).Save(task).Error
}

func (r TasksRepository) Create(ctx context.Context, task *entity.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r TasksRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(entity.Task{}, id).Error
}
