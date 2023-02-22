package model

type Task struct {
	ID             uint64 `gorm:"primaryKey"`
	SequenceID     uint64
	Name           string   `json:"title"`
	Description    string   `json:"description"`
	Tests          string   `json:"tests"`
	AuthorSolution Solution `gorm:"foreignKey:ID"`
	Tags           []Tag    `gorm:"foreignKey:ID"`
}
