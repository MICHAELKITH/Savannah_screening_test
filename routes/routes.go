package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savannah/sms/controllers"
)

func SetupRoutes(app *fiber.App) {

    app.Post("/customers", controllers.AddCustomer)
    app.Post("/orders", controllers.AddOrder)
	app.Put("/orders/:id", controllers.AddOrder)
}


