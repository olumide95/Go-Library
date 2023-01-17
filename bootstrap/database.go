package bootstrap

import (
	"log"
	"os"

	"github.com/olumide95/go-library/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) InitDb() *Database {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{}, &models.Book{}, &models.BookLog{})

	d.DB = db

	return d
}
