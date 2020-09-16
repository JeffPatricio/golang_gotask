package main

import (
	"fmt"
	"gotask/database"
	"gotask/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	port := os.Getenv("PORT")
	database.TestConnection()
	fmt.Printf("Api running on port %s\n", port)
	r := routes.GetRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
