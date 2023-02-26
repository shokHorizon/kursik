package model

type Solution struct {
	ID     uint64 `gorm:"primaryKey"`
	UserID uint64 `gorm:"default:NULL"`
	TaskID uint64 `gorm:"default:NULL"`
	Code   string `json:"code"`
	Status bool   `gorm:"default:null"`
}
