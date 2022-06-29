package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/caiostarke/restApi-and-grpc/app/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateNewAccessToken(u models.UserResponse) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	minuteCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minuteCount)).Unix()
	claims["user"] = u.Username
	claims["role"] = u.Role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
