package database

import (
	"api_test/config"
	"api_test/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	var err error
	//postgres://postgres:sun@127.0.0.1:5432/
	// open connection, config DATABASE_URL is dsn
	db_str := config.Db_url + config.Db_usr + ":" + config.Db_pwd + "@" + config.Db_host + ":" + config.Db_port + "/" + config.Db_name
	DB, err = gorm.Open(postgres.Open(db_str), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database:", err.Error())
	} else {
		fmt.Println("Initialized connection to databse")
	}

	// auto create table in db
	DB.AutoMigrate(&model.Staff{})

}
