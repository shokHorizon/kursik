package postgres

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewPostgres(repository Config) (*Repository, error) {
	var err error
	db, err := gorm.Open(mysql.Open(repository.ConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
		return nil, err
	}
	return &Repository{db: db}, nil
}
