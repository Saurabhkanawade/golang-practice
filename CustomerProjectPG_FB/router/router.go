package router

import (
	"github.com/gofiber/fiber"
	"github.com/saurabhkanawade/customerprojectpg_fb/service"
)

func Router() *fiber.App {
	route := fiber.New()

	route.Get("/customers", service.GetCustomers)
	route.Get("/customers/:id", service.GetCustomer)
	route.Put("/customers/:id", service.UpdateCustomer)
	route.Post("/customers", service.CreateCustomer)
	route.Delete("/customers/:id", service.DeleteCustomer)

	return route
}
