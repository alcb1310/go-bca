package models

import (
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/0x4149/logz"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var err error

type DB struct {
	Data *gorm.DB
}

func (d *DB) Initialize() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, dbPort)

	d.Data, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logz.Fatal("Unable to connect to the database")
	}

	d.Data.AutoMigrate(&Company{})
	d.Data.AutoMigrate(&User{})
	d.Data.AutoMigrate(&Project{})
	d.Data.AutoMigrate(&Supplier{})
	d.Data.AutoMigrate(&BudgetItem{})
	d.Data.AutoMigrate(&Budget{})
	d.Data.AutoMigrate(&Invoice{})
	d.Data.AutoMigrate(&InvoiceDetails{})
	d.Data.AutoMigrate(&Historic{})
	d.Data.AutoMigrate(&LoggedInUser{})

	logz.Info("Database connected")
}
