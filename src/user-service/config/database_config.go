package config

import (
	"chilindo/src/user-service/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	host     string = "localhost"
	port     string = "3306"
	username string = "root"
	password string = "A1231230"
	database string = "chilindo"
)
var connectString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	username,
	password,
	host,
	port,
	database,
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	//if err := godotenv.Load("local.env"); err != nil {
	//	panic("Error loading .env.admin file")
	//}
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASS"),
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_PORT"),
	//	os.Getenv("DB_NAME"),
	//)
	DB, err = gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	if errConnect := DB.AutoMigrate(&entity.User{}, &entity.Address{}); errConnect != nil {
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
