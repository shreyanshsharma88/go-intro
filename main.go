package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-api-test/handlers"
	"go-api-test/middleware"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// http.HandleFunc("/ ", rootHandler)
	router.HandleFunc("/users", middleware.CORSMiddleware(handlers.UserHandler))
	router.HandleFunc("/user/signUp", middleware.CORSMiddleware(handlers.SignUpHandler))
	router.HandleFunc("/user/login", middleware.CORSMiddleware(handlers.LoginHandler))
	router.HandleFunc("/users/todo", middleware.CORSMiddleware(middleware.TodoMiddleware(handlers.TodoHandler)))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
