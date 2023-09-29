package controller

import (
	"github.com/gofiber/fiber"
	"github.com/mohdaalam005/fiber/services"
)

func Route(route *fiber.App) {
	route.Get("/students", services.GetStudents)
	route.Get("/students/:id", services.GetStudent)
	route.Delete("/students/:id", services.DeleteStudent)
	route.Post("/students", services.CreateStudent)
	route.Put("/students/:id", services.UpdateStudent)
	route.Listen(":8080")

}
