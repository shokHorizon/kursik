package model

type Solution struct {
	ID     uint64 `gorm:"primaryKey"`
	UserID uint64 `gorm:"foreignKey:ID"`
	TaskID uint64 `gorm:"foreignKey:ID"`
	Code   string `json:"description"`
	IsDone bool
}
