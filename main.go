package main

import (
	"fmt"
	"go-api-test/handlers"
	"net/http"
	"github.com/gorilla/mux"
	// "go-api-test/middleware"
)


func main() {
	router := mux.NewRouter()

	// http.HandleFunc("/ ", rootHandler)
	router.HandleFunc("/users", handlers.UserHandler)
	router.HandleFunc("/user/signUp", handlers.SignUpHandler)
	router.HandleFunc("/user/login", handlers.LoginHandler)
	router.HandleFunc("/users/todo", handlers.TodoHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
