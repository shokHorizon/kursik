package model

type User struct {
	ID             uint64     `gorm:"primaryKey"`
	Login          string     `json:"login"`
	HashedPassword uint64     `json:"hashedPassword"`
	AccessLevel    uint16     `json:"accessLevel"`
	Tasks          []Task     `gorm:"foreignKey:ID"`
	Solutions      []Solution `gorm:"foreignKey:ID"`
}
