package utils

import (
	"go-api-test/models"
	"go-api-test/store"
)

func FindUserByName(name string) (models.User, int) {
	for i, user := range store.Data {
		if user.Name == name {
			return user, i
		}
	}
	return models.User{}, -1
}

func FindUserByID(id string) (models.User, int) {
	for i, user := range store.Data {
		if user.Id == id {
			return user, i
		}
	}
	return models.User{}, -1
}
