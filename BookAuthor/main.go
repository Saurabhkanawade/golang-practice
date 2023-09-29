package main

import (
	"github.com/SaurabhKanawade/BookAuthor/controller"
	"github.com/SaurabhKanawade/BookAuthor/router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the server on localhost 9000....................")

	// loading the app properties file
	err := godotenv.Load("app.env")
	controller.CheckNilError(err)

	//calling to the router file
	route := router.Router()

	err = http.ListenAndServe(":9000", route)
	controller.CheckNilError(err)

}
