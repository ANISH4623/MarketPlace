package database

import (
	"fmt"
	"log"

	"github.com/ANISH4623/helpers"
	"github.com/ANISH4623/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		helpers.AppConfig.DB_USER, helpers.AppConfig.DB_PASSWORD, helpers.AppConfig.DB_HOST, helpers.AppConfig.DB_PORT,
		helpers.AppConfig.DB_NAME)
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
	log.Println("Connected to the database successfully")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	db.AutoMigrate(&models.User{}, &models.Product{})

	Database = DbInstance{Db: db}
}
