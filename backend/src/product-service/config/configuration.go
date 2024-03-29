package config

import (
	"backend/src/product-service/entity"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	const (
		env = ".env"
	)
	if err := godotenv.Load(env); err != nil {
		log.Println("Error loading .env in product file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err.Error())
	}
	if errConnect := DB.AutoMigrate(&entity.Product{}, &entity.ProductOption{}, &entity.ProductImage{}); errConnect != nil {
		panic(errConnect.Error())
	}
}

func GetDB() *gorm.DB {
	return DB
}

func init() {
	ConnectDatabase()
	log.Println("Connected to database...")
}

func CloseDatabase(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection")
	}
	dbSQL.Close()
}
