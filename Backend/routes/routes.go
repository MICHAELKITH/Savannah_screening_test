package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savannah/sms/controllers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/api/customers", controllers.GetCustomers)
	app.Post("/api/customers", controllers.AddCustomer)
	app.Get("/api/orders", controllers.GetOrders)
	app.Post("/api/orders", controllers.AddOrder)
	
}


