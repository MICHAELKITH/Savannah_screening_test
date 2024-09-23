package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savannah/sms/controllers" // Adjust based on your project structure
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go Fiber with Okta OpenID Connect!!")
	})

	app.Get("/login", controllers.Login)
	app.Get("/callback", controllers.Callback)
	app.Post("/users", controllers.CreateUser)

	// Protected routes
	app.Post("/customers", controllers.Protect, controllers.AddCustomer)
	app.Post("/orders", controllers.Protect, controllers.AddOrder)
	app.Put("/orders/:id", controllers.Protect, controllers.AddOrder)
}
