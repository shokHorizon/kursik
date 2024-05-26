package entity

import "github.com/oapi-codegen/runtime/types"

type Task struct {
	ID          types.UUID `gorm:"column:id;primary_key"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	IsApproved  bool       `gorm:"column:is_approved"`
	UserID      int64      `gorm:"column:user_id"`
}

type Tasks = []*Task

type TaskFilter struct {
	ID          *string
	Name        *StringFilter
	Description *StringFilter
	IsApproved  bool
	UserID      *IntFilter
}
