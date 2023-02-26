package model

type Task struct {
	ID             uint64 `gorm:"primaryKey"`
	SequenceID     uint64
	Name           string   `json:"title"`
	Description    string   `json:"description"`
	Tests          string   `gorm:"default:NULL"`
	AuthorSolution Solution `gorm:"default:NULL"`
	Tags           []Tag    `gorm:"default:NULL"`
	CourseId       uint64   `gorm:"default:NULL"`
	UserId         uint64   `gorm:"default:NULL" json:"user_id"`
}
