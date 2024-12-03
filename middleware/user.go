package middleware

import (
	"context"
	"go-api-test/utils"
	"net/http"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret")

func UserAuthMiddleware(req *http.Request, res http.ResponseWriter, next http.HandlerFunc) {
	tokenString := req.Header.Get("token")
	if tokenString == "" {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userId := token.Claims.(jwt.MapClaims)["id"]
	user, _ := utils.FindUserByID(userId.(string))

	// add user to context
	ctx := context.WithValue(req.Context(), "user", user)

	next(res, req.WithContext(ctx))

}
  