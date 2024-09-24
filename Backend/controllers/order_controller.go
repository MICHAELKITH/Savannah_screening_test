package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/savannah/sms/config"
	"github.com/savannah/sms/services"
	"github.com/savannah/sms/clients"
)

type Order struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Item       string  `json:"item"`
	Amount     float64 `json:"amount"`
}

// AddOrder handles the addition of a new order
func AddOrder(c *fiber.Ctx) error {
	order := new(Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	// Check if customer exists
	var exists bool
	err := config.DBPool.QueryRow(context.Background(), "SELECT EXISTS (SELECT 1 FROM customers WHERE id=$1)", order.CustomerID).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if customer exists: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error checking customer")
	}
	if !exists {
		log.Printf("Customer ID %d not found\n", order.CustomerID)
		return c.Status(fiber.StatusBadRequest).SendString("Customer not found")
	}

	// Insert the order into the database
	query := `INSERT INTO orders (customer_id, item, amount) VALUES ($1, $2, $3)`
	_, err = config.DBPool.Exec(context.Background(), query, order.CustomerID, order.Item, order.Amount)
	if err != nil {
		log.Printf("Unable to insert order: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add order")
	}

	// Retrieve customer details for SMS notification
	var customerName, customerCode string
	err = config.DBPool.QueryRow(context.Background(), "SELECT name, code FROM customers WHERE id=$1", order.CustomerID).Scan(&customerName, &customerCode)
	if err != nil {
		log.Printf("Error retrieving customer details: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving customer details")
	}

	// Send SMS notification using Africa's Talking
	message := fmt.Sprintf("Hi %s, your order for %s has been received. Amount: %.2f", customerName, order.Item, order.Amount)
	err = sendSMS("+254714707147", message)
	if err != nil {
		log.Printf("Unable to send SMS: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Order added but failed to send SMS")
	}

	return c.Status(fiber.StatusOK).SendString("Order added successfully and SMS sent!")
}

// GetOrders retrieves all orders from the database
func GetOrders(c *fiber.Ctx) error {
	rows, err := config.DBPool.Query(context.Background(), "SELECT id, customer_id, item, amount FROM orders")
	if err != nil {
		log.Printf("Error fetching orders: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch orders")
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.Item, &order.Amount); err != nil {
			log.Printf("Error scanning order: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to parse orders")
		}
		orders = append(orders, order)
	}

	return c.JSON(orders) //  orders as JSON
}

// Talking service
func sendSMS(phoneNumber, message string) error {
	// Get credentials from environment variables
	atUsername := os.Getenv("AT_USERNAME")
	atAPIKey := os.Getenv("AT_API_KEY")
	env := os.Getenv("ENV") 

	client, errClient := clients.NewDefaultHttpClient(&http.Client{})
	if errClient != nil {
		fmt.Println("error client ", errClient.Error())
	}
	// Initialize SMS service
	smsSrv, errSms := services.NewSmsService(atUsername, atAPIKey, env, client)
	if errSms != nil {
		fmt.Println("error client ", errSms.Error())
	}

	// Call send
	str, errStr := smsSrv.Send("", phoneNumber, message)
	if errStr != nil {
		fmt.Println("error client ", errStr.Error())
	}
	log.Printf("SMS sent: %v\n", str)
	return nil
}
