package entity

// Course defines model for Course with gorm annotations
type Course struct {
	ID          int64  `json:"id" gorm:"column:id;primary_key"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	IsApproved  bool   `json:"is_approved" gorm:"column:is_approved"`
	UserID      int64  `json:"user_id" gorm:"column:user_id"`
}

type Courses = []*Course
