package token

import (
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var jwtKey []byte

type IJwtMiddleware interface {
	GenerateJWT(email string, userid uint) (tokenString string, err error)
	ExtractToken(tokenString string) (*Claims, error)
}

type Claims struct {
	Email    string
	UserId   uint
	Username string
	jwt.StandardClaims
}

func (j *Claims) GenerateJWT(email string, userid uint, userName string) (tokenString string, err error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env in jwt-token file")
	}
	jwtKey = []byte(os.Getenv("SECRET_KEY"))

	expirationTime := time.Now().Add(24 * 7 * time.Hour)
	claims := &Claims{
		Email:    email,
		UserId:   userid,
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ExtractToken(signedToken string) (*Claims, error) {
	jwtKey = []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
