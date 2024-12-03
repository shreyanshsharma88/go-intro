package auth

import (
	"time"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret")

// GenerateJWT creates a JWT token for a user
func GenerateJWT(name string, id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["name"] = name
	claims["id"] = id
	tokenString, _ := token.SignedString(secretKey)
	return tokenString
}
