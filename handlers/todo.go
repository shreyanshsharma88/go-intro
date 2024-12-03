package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"go-api-test/models"
	"go-api-test/utils"
)

func TodoHandler(res http.ResponseWriter, req *http.Request) {
	user := req.Context().Value("user").(models.User)
	

	switch req.Method {
	case http.MethodGet:
		utils.SendJsonResponse(res, http.StatusOK, user.Todos)

	case http.MethodPost:
		var newTodo models.Todo
		if err := json.NewDecoder(req.Body).Decode(&newTodo); err != nil {
			http.Error(res, "Invalid request body", http.StatusBadRequest)
			return
		}
		newTodo.ID = uuid.New().String()
		
		user.Todos = append(user.Todos, newTodo)

		utils.SendJsonResponse(res, http.StatusCreated, newTodo)

	case http.MethodDelete:
		todoId := req.URL.Query().Get("todoId")
		_, todoIndex := utils.FindUserTodoByID(todoId, &user)
		if todoIndex == -1 {
			http.Error(res, "Todo not found", http.StatusNotFound)
			return
		}
		// store.Data[userIndex].Todos = append(store.Data[userIndex].Todos[:todoIndex], store.Data[userIndex].Todos[todoIndex+1:]...)
		user.Todos = append(user.Todos[:todoIndex], user.Todos[todoIndex+1:]...)
		utils.SendJsonResponse(res, http.StatusOK, user.Todos)

	case http.MethodPut:
		todoId := req.URL.Query().Get("todoId")
		var updatedTodo models.Todo
		if err := json.NewDecoder(req.Body).Decode(&updatedTodo); err != nil {
			http.Error(res, "Invalid request body", http.StatusBadRequest)
			return
		}
		todo, todoIndex := utils.FindUserTodoByID(todoId, &user)
		if todoIndex == -1 {
			http.Error(res, "Todo not found", http.StatusNotFound)
			return
		}
		todo.Title = updatedTodo.Title
		todo.Priority = updatedTodo.Priority
		// store.Data[userIndex].Todos[todoIndex] = todo
		user.Todos[todoIndex] = todo
		utils.SendJsonResponse(res, http.StatusOK, user)

	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}