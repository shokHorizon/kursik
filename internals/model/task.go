package model

type Task struct {
	ID          uint64 `gorm:"primaryKey"`
	SequenceID  uint64 `gorm:"default:NULL" json:"sequence_id"`
	Name        string `json:"title"`
	Description string `json:"description"`
	Tests       string `json:"tests"`
	Tags        []Tag  `gorm:"many2many:tasks_tags" json:"tags"`
	CourseID    uint64 `gorm:"default:NULL" json:"course_id"`
	UserID      uint64 `gorm:"default:NULL" json:"user_id"`
}
