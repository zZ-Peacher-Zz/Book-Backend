package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims = claims
	tk, _ := token.SignedString([]byte("mysecretkey"))
	return tk
}
