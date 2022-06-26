package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type variables struct {
	DBPassword string
	DBUser     string
	DBHost     string
	DBPort     int
	DBName     string
}

var Env variables

func init() {
	err := godotenv.Load()
	var dbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln("Unable to load env variables", err)
	}
	Env = variables{
		DBPassword: os.Getenv("DB_PASS"),
		DBUser:     os.Getenv("DB_USER"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort: func() int {
			if dbPort == 0 {
				return 5432
			} else {
				return dbPort
			}
		}(),
		DBName: os.Getenv("DB_NAME"),
	}
}
