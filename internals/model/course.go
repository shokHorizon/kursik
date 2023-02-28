package model

type Course struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Tasks   []Task `json:"tasks"`
	OwnerID uint64 `json:"owner_id"`
	Owner   User   `gorm:"foreignKey: OwnerID"`
}
