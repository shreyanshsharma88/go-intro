package middleware

import (
	"context"
	"go-api-test/utils"
	"net/http"
	"github.com/golang-jwt/jwt"
)

func TodoMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("token")
		if token == "" {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

			return []byte("secret"), nil
		})


		if err != nil || !jwtToken.Valid {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}
		userId := jwtToken.Claims.(jwt.MapClaims)["id"]
		user, _ := utils.FindUserByID(userId.(string))

		todoId := req.URL.Query().Get("todoId")
		if todoId != "" {
			_, todoIndex := utils.FindUserTodoByID(todoId, &user)
			if todoIndex == -1 {
				http.Error(res, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		ctx := context.WithValue(req.Context(), "user", user)
		next(res, req.WithContext(ctx))

	}
}
