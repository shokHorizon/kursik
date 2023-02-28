package model

type Solution struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	UserID uint64 `json:"user_id"`
	TaskID uint64 `json:"task_id"`
	Code   string `json:"code"`
	Status bool   `gorm:"default:null" json:"status"`
}
