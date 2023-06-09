package database

import (
	"api_test/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	var err error

	// open connection, config DATABASE_URL is dsn
	db_str := config.Db_url + config.Db_name
	DB, err = gorm.Open(postgres.Open(db_str), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database:", err.Error())
	} else {
		fmt.Println("Initialized connection to databse")
	}

}
