package database

import (
	"blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DBConn *gorm.DB

func ConnectDB() {

	host := os.Getenv("db_host")
	dbUser := os.Getenv("db_user")
	dbName := os.Getenv("db_name")

	dsn := dbUser + ":@tcp(" + host + ":3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

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
