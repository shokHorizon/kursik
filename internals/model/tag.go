package model

type Tag struct {
	ID   uint64 `gorm:"unique"`
	Name string `gorm:"primaryKey"`
}
