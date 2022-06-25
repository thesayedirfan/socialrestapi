package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "host=db user=dbuser password=pass123 dbname=db port=5432 sslmode=disable"
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

}
