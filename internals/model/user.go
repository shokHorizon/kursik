package model

import "fmt"

type User struct {
	ID             uint64 `gorm:"primaryKey"`
	Login          string `gorm:"unique" json:"login" form:"login"`
	HashedPassword string `json:"hashedPassword" form:"password"`
	AccessLevel    uint16 `json:"accessLevel"`
	Tasks          []Task
	Solutions      []Solution
}

var users map[string]string

func init() {
	// Mock for db interaction
	users = map[string]string{
		"admin": "123",
		"bibka": "123",
	}
}

func (user *User) LoginUser() bool {
	if valid_password, ok := users[user.Login]; !ok || valid_password != user.HashedPassword {
		fmt.Println("Login failed for", user.Login)
		return false
	}
	return true
}
