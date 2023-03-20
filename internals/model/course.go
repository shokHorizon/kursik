package model

type Course struct {
	ID      uint64  `gorm:"primaryKey" json:"id"`
	Title   string  `json:"title"`
	Tasks   []Task  `json:"tasks"`
	OwnerID uint64  `json:"owner_id" gorm:"default:NULL"`
	Owner   User    `gorm:"foreignKey: OwnerID"`
	Users   []*User `gorm:"many2many:courses_users"`
}
