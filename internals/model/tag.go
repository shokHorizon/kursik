package model

type Tag struct {
	ID    uint64 `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"unique" json:"name"`
	Tasks []Task `gorm:"many2many:tasks_tags" json:"tasks"`
}
