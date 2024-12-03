package store

import "go-api-test/models"

// Global data variable to store users
var Data = []models.User{
	{
		Name:     "Shreyansh",
		Password: "123",
		Id:       "1",
		Todos: []models.Todo{
			{ID: "1", Title: "Buy Milk", Priority: "High"},
		},
	},
}