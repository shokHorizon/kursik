package course

import (
	"github.com/shokHorizon/kursik/internals/models/task"
)

type Course struct {
	ID    uint64 `gorm:"primaryKey"`
	Title string `json:"title"`
	Tasks []task.Task
}
