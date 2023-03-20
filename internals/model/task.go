// package model

// type Task struct {
// 	ID          uint64 `gorm:"primaryKey" json:"id"`
// 	SequenceID  uint64 `gorm:"default:NULL" json:"sequence_id"`
// 	Name        string `json:"title"`
// 	Description string `json:"description"`
// 	Tests       string `json:"tests"`
// 	Tags        []Tag  `gorm:"foreignKey:TaskID"`
// 	CourseID    uint64 `gorm:"default:NULL" json:"course_id"`
// 	UserID      uint64 `json:"user_id"`
// }

package model

type Task struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	SequenceID  uint64 `gorm:"default:NULL" json:"sequence_id"`
	Name        string `json:"title"`
	Description string `json:"description"`
	Tests       string `json:"tests"`
	Tags        []*Tag `gorm:"many2many:tasks_tags"`
	CourseID    uint64 `gorm:"default:NULL" json:"course_id"`
	UserID      uint64 `json:"user_id"`
}
