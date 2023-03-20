package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID             uint64     `gorm:"primaryKey" json:"id"`
	Login          string     `gorm:"unique" json:"login" form:"login"`
	HashedPassword string     `json:"password" form:"password"`
	AccessLevel    uint16     `json:"accessLevel"`
	Tasks          []Task     `json:"tasks"`
	Solutions      []Solution `json:"solutions"`
	Courses        []*Course  `gorm:"many2many:courses_users"`
	Email          string     `json:"email"`
}

type SignUpInput struct {
	Login           string `json:"login" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	AccessLevel     uint16 `json:"accessLevel"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

// type UserResponse struct {
// 	ID          uint64 `json:"id"`
// 	Login       string `json:"login"`
// 	Email       string `json:"email"`
// 	AccessLevel uint16 `json:"accessLevel"`
// }

// func FilterUserRecord(user *User) UserResponse {
// 	return UserResponse{
// 		ID:          user.ID,
// 		Login:       user.Login,
// 		Email:       user.Email,
// 		AccessLevel: user.AccessLevel,
// 	}
// }

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
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
