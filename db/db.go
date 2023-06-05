package db

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		dsn := os.Getenv("DB_URL")
		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Error connecting to the database")
		}
		db = conn
	})
	return db
}
