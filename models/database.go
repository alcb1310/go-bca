package models

import (
	"fmt"
	"os"

	"gitlab.com/0x4149/logz"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	data *gorm.DB
}

func Initialize() DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logz.Fatal("Unable to connect to the database")
	}

	return DB{data: db}
}
