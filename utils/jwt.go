package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secret = "sldkfjsdlk78gdlfkjsdflksdjfldskjl4j5908rjwedoif"

func GenerateToken(userId int64, email string, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"name":    name,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secret))
}
