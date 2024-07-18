package database

import (
	"blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Successfully connected to database")

	db.AutoMigrate(new(model.Blog))
	DBConn = db
}
