package task

import (
	"github.com/shokHorizon/kursik/internals/models/tag"
)

type Task struct {
	ID             uint64 `gorm:"primaryKey"`
	SequenceID     uint64
	Name           string `json:"title"`
	Description    string `json:"description"`
	Tests          string `json:"tests"`
	AuthorSolution uint64
	Tags           []tag.Tag
}
