package db

import (
	"fmt"
	"log"
	"todo/config"
	"todo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDbUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.Env.DBUser,
		config.Env.DBPassword,
		config.Env.DBHost,
		config.Env.DBPort,
		config.Env.DBName,
	)
}

func Start() {
	// Creating a connection to the database
	var err error
	log.Println("db url", GetDbUrl())
	DB, err = gorm.Open(postgres.Open(GetDbUrl()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// run the migrations: todo struct
	DB.AutoMigrate(&models.Todo{})
}
