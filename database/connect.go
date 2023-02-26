package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/shokHorizon/kursik/config"
	"github.com/shokHorizon/kursik/internals/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Wrong db port")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	DB.Migrator().DropTable(&model.User{}, &model.Tech{}, &model.Course{}, &model.Tag{}, &model.Task{}, &model.Solution{})
	DB.AutoMigrate(&model.User{}, &model.Tech{}, &model.Course{}, &model.Tag{}, &model.Task{}, &model.Solution{})
	fmt.Println("Database Migrated")
}
