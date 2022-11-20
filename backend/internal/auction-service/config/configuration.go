package config

import (
	"backend/internal/auction-service/entity"
	"backend/pkg/websocket"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env in auction file")
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
		panic(err.Error())
	}
	if errConnect := DB.AutoMigrate(&entity.Auction{}, &entity.Bid{}); errConnect != nil {
		panic(errConnect.Error())
	}
}

func GetDB() *gorm.DB {
	return DB
}

var (
	Pool *websocket.Pool
)

func init() {
	Pool = websocket.NewPool()
	go Pool.Start()
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
