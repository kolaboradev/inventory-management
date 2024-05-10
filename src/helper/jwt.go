package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	staffentity "github.com/kolaboradev/inventory/src/models/entities/staff"
)

func GenerateTokenJWT(staff staffentity.Staff) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["staffId"] = staff.Id
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	claims["phoneNumber"] = staff.PhoneNumber
	claims["roles"] = "staff"

	secretKeys := os.Getenv("JWT_SECRET")

	secretToken, err := token.SignedString([]byte(secretKeys))
	ErrorIfPanic(err)

	return secretToken
}

func CheckTokenJWT(t *jwt.Token) (interface{}, error) {
	secretJWT := os.Getenv("JWT_SECRET")
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method")
	}

	return []byte(secretJWT), nil
}
