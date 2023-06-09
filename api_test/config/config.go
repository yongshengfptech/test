package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Env string
var Db_url string
var Host string
var Db_name string

func Load_env() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("load env failed")
	} else {
		fmt.Println("load env suceed")
	}

	Env = os.Getenv("ENVIRONMENT")
	Db_url = os.Getenv("DATABASE_URL")
	Host = os.Getenv("HOST")
	Db_name = os.Getenv("DATABASE_NAME")

	fmt.Println(Env, " / ", Db_url, " / ", Host, " / ", Db_name)
}
