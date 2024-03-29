package config

import (
	"backend/src/account-service/entity"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
)

var (
	DB          *gorm.DB
	err         error
	CookieAuth  string
	CookieStore cookie.Store
	NewSessions []string
)

func ConnectDatabase() {
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
	if errConnect := DB.AutoMigrate(&entity.Account{}, &entity.Address{}); errConnect != nil {
		panic(errConnect.Error())
	}
}

func GetDB() *gorm.DB {
	return DB
}

func init() {
	const (
		env = ".env"
	)
	if err := godotenv.Load(env); err != nil {
		log.Println("Error loading .env in account-config file")
	}

	//Create store for sessions
	keyPairs := []byte(os.Getenv("KEY_PAIRS"))
	CookieStore = cookie.NewStore(keyPairs)
	CookieStore.Options(sessions.Options{
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: true, // prevent from FE to be accessed from client using script
		Secure:   true, //use in https. It means enable this line when deploying the project
		SameSite: http.SameSiteNoneMode,
	}) // 7 days

	//List of cookies name
	CookieAuth = os.Getenv("COOKIE_AUTH")

	//Create List of cookies
	//NewSessions = append(NewSessions, CookieAuth)

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
