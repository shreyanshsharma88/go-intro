package handlers

import (
	"encoding/json"
	"net/http"
	"go-api-test/auth"
	"go-api-test/models"
	"go-api-test/utils"
	"go-api-test/store"
	"github.com/google/uuid"
)



func UserHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		var newUser models.User
		if err := json.NewDecoder(req.Body).Decode(&newUser); err != nil {
			http.Error(res, "Invalid request body", http.StatusBadRequest)
			return
		}
		store.Data = append(store.Data, newUser)
		utils.SendJsonResponse(res, http.StatusCreated, newUser)
		return
	}
	http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
}

func SignUpHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		var user models.User
		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			http.Error(res, "Invalid request body", http.StatusBadRequest)
		}
		_, existingUserIndex := utils.FindUserByName(user.Name)
		if existingUserIndex == -1 {
			user.Id = uuid.New().String()
			store.Data = append(store.Data, user)
			token := auth.GenerateJWT(user.Name, user.Id)
			utils.SendJsonResponse(res, http.StatusCreated, map[string]string{"token": token})
			return
		}
		http.Error(res, "Username exists", http.StatusConflict)
		return
	}
	http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		var details models.User
		if err := json.NewDecoder(req.Body).Decode(&details); err != nil {
			http.Error(res, "Invalid request body", http.StatusBadRequest)
		}
		user, existingUserIndex := utils.FindUserByName(details.Name)
		if existingUserIndex == -1 {
			http.Error(res, "User not found", http.StatusNotFound)
			return
		}
		if user.Password == details.Password {
			token := auth.GenerateJWT(user.Name, user.Id)
			utils.SendJsonResponse(res, http.StatusOK, map[string]string{"token": token})
			return
		}
		http.Error(res, "Invalid password", http.StatusUnauthorized)
	}
	http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
}
