package model

type Course struct {
	ID    uint64 `gorm:"primaryKey"`
	Title string `json:"title"`
	Tasks []Task `gorm:"foreignKey:ID"`
}
