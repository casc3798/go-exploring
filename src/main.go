package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/casc3798/go-exploring/src/api/users"
	"github.com/casc3798/go-exploring/src/config/database"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	db, err := database.NewDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
	defer db.Close()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)

	router.Mount("/users", users.Router())

	fmt.Printf("Starting server on port %v\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), router))
}
