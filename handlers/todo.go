package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"go-api-test/models"
	"go-api-test/store"
	"go-api-test/utils"
)

func TodoHandler(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	user, userIndex := utils.FindUserByName(name)
	if userIndex == -1 {
		http.Error(res, "User not found", http.StatusNotFound)
		return
	}

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
		store.Data[userIndex].Todos = append(store.Data[userIndex].Todos, newTodo)
		utils.SendJsonResponse(res, http.StatusCreated, newTodo)

	case http.MethodDelete:
		todoId := req.URL.Query().Get("todoId")
		_, todoIndex := utils.FindUserTodoByID(todoId, &store.Data[userIndex])
		if todoIndex == -1 {
			http.Error(res, "Todo not found", http.StatusNotFound)
			return
		}
		store.Data[userIndex].Todos = append(store.Data[userIndex].Todos[:todoIndex], store.Data[userIndex].Todos[todoIndex+1:]...)
		utils.SendJsonResponse(res, http.StatusOK, store.Data[userIndex].Todos)

	case http.MethodPut:
		todoId := req.URL.Query().Get("todoId")
		var updatedTodo models.Todo
		if err := json.NewDecoder(req.Body).Decode(&updatedTodo); err != nil {
			http.Error(res, "Invalid request body", http.StatusBadRequest)
			return
		}
		todo, todoIndex := utils.FindUserTodoByID(todoId, &store.Data[userIndex])
		if todoIndex == -1 {
			http.Error(res, "Todo not found", http.StatusNotFound)
			return
		}
		todo.Title = updatedTodo.Title
		todo.Priority = updatedTodo.Priority
		store.Data[userIndex].Todos[todoIndex] = todo
		utils.SendJsonResponse(res, http.StatusOK, store.Data[userIndex])

	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}