package user

import (
	"github.com/shokHorizon/kursik/internals/models/solution"
	"github.com/shokHorizon/kursik/internals/models/task"
)

type User struct {
	ID             uint64 `gorm:"primaryKey"`
	Login          string `json:"login"`
	HashedPassword uint64 `json:"hashedPassword"`
	AccessLevel    uint16 `json:"accessLevel"`
	Tasks          []task.Task
	Solutions      []solution.Solution
}
