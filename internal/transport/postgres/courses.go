package postgres

import (
	"context"
	"github.com/shokHorizon/kursik/internal/entity"
	service "github.com/shokHorizon/kursik/internal/serivce"
	"gorm.io/gorm"
)

var _ service.Courses = (*CoursesRepository)(nil)

type CoursesRepository struct {
	db *gorm.DB
}

func NewCoursesRepository(db *gorm.DB) *CoursesRepository {
	return &CoursesRepository{db: db}
}

func (r CoursesRepository) Get(ctx context.Context, id int64) (*entity.Course, error) {
	var course entity.Course
	err := r.db.WithContext(ctx).First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r CoursesRepository) GetAll(ctx context.Context, params entity.Pagination) (entity.Courses, error) {
	var courses entity.Courses
	err := r.db.WithContext(ctx).Limit(int(*params.Limit)).Offset(int(*params.Offset)).Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (r CoursesRepository) Search(ctx context.Context, params entity.SearchParams) (entity.Courses, error) {
	var courses entity.Courses
	err := r.db.WithContext(ctx).
		Where("name LIKE?", "%"+*params.Search+"%").Limit(int(*params.Limit)).Offset(int(*params.Offset)).
		Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (r CoursesRepository) Patch(ctx context.Context, course *entity.Course) error {
	return r.db.WithContext(ctx).Save(course).Error
}

func (r CoursesRepository) Create(ctx context.Context, course *entity.Course) error {
	return r.db.WithContext(ctx).Create(course).Error
}

func (r CoursesRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(entity.Course{}, id).Error
}
