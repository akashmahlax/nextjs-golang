package main

import (
	"log"
	"net/http"
	"nextjs-golang/internal/db"
	"nextjs-golang/internal/handlers"
	"nextjs-golang/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect("mongodb://localhost:27017")

	r := mux.NewRouter()

	r.HandleFunc("/api/auth/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/api/auth/login", handlers.Login).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.Authenticate)
	api.HandleFunc("/services", handlers.GetServices).Methods("GET")
	api.HandleFunc("/services", handlers.CreateService).Methods("POST")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
