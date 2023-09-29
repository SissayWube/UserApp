package main

import (
	db "UserApp/database"
	"UserApp/handler"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading env file")
	}
	// Initialize the database
	err = db.InitDB()
	if err != nil {
		panic(err)
	}

	dbc, err := db.GetDBCon().DB()
	if err != nil {
		panic(err)
	}

	defer dbc.Close()

	// Start server
	handler.StartServer()

}
