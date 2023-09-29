package main

import (
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/mohdaalam005/fiber/controller"
	"github.com/mohdaalam005/fiber/database"
	"log"
)

func main() {
	app := fiber.New()
	log.Println("application has been started ")
	godotenv.Load(".env")
	database.Connect()
	controller.Route(app)

}
