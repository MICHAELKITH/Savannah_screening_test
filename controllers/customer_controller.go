package controllers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/savannah/sms/config"
)

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// AddCustomer adds a new customer to the database
func AddCustomer(c *fiber.Ctx) error {
	customer := new(Customer)
	if err := c.BodyParser(customer); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	query := `INSERT INTO customers (name, code) VALUES ($1, $2) RETURNING id`
	err := config.DBPool.QueryRow(context.Background(), query, customer.Name, customer.Code).Scan(&customer.ID)
	if err != nil {
		log.Printf("Unable to insert customer: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add customer")
	}

	return c.Status(fiber.StatusOK).SendString("Customer added successfully!")
}


// Business logic
// Http // api layer or data layer
// deploy