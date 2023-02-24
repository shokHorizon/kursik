package model

type Tech struct {
	ID    uint64 `gorm:"primaryKey"`
	Title string `gorm:"unique" json:"title"`
}
