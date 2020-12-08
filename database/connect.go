package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgre() (*gorm.DB, error) {


	//conf := config.Get()

	db, err := gorm.Open(postgres.Open("host=192.168.25.130 user=postgres password=123456 dbname=crudapi port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})

	if err == nil {
		DB = db
		err := db.AutoMigrate()
		return db, err
	}
	return nil, err
	
}
