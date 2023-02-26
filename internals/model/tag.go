package model

type Tag struct {
	ID    uint64 `gorm:"primaryKey"`
	Name  string `gorm:"unique" json:"title"`
	Tasks []Task `gorm:"many2many:tasks_tags"`
}
