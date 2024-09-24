package controllers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/savannah/sms/config"
)

//  structure of a customer
type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// adds a new customer to the database
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

	return c.Status(fiber.StatusOK).JSON(customer)
}

// GetCustomers fetches all customers from the database
func GetCustomers(c *fiber.Ctx) error {
	rows, err := config.DBPool.Query(context.Background(), "SELECT id, name, code FROM customers")
	if err != nil {
		log.Printf("Error fetching customers: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch customers")
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Code); err != nil {
			log.Printf("Error scanning customer: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch customers")
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating customers: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch customers")
	}

	return c.Status(fiber.StatusOK).JSON(customers)
}
