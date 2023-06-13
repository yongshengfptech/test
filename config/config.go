package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Env string
var Host string

var Db_url string
var Db_usr string
var Db_pwd string
var Db_name string
var Db_host string
var Db_port string

func Load_env() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("load env failed")
	} else {
		fmt.Println("load env suceed")
	}

	Env = os.Getenv("ENVIRONMENT")
	Host = os.Getenv("HOST")

	Db_url = os.Getenv("DATABASE_URL")
	Db_name = os.Getenv("DATABASE_NAME")
	Db_usr = os.Getenv("DATABASE_USER")
	Db_pwd = os.Getenv("DATABASE_PASSWORD")
	Db_host = os.Getenv("DATABASE_HOST")
	Db_port = os.Getenv("DATABASE_PORT")

	fmt.Println(Env, " / ", Db_url, " / ", Host, " / ", Db_name)
}
