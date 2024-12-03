package utils
import (
	"go-api-test/models"

)

func FindUserTodoByID(todoId string, user *models.User) (models.Todo, int) {
	for i, todo := range user.Todos {
		if todo.ID == todoId {
			return todo, i
		}
	}
	return models.Todo{}, -1
}