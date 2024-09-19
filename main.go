package main

import (
	"log"

	"net/http"

	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/savannah/sms/clients"
	"github.com/savannah/sms/config"
	"github.com/savannah/sms/routes"
	"github.com/savannah/sms/services"
)

func main() {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return
	}
	// Get credentials from environment variables
	atUsername := os.Getenv("AT_USERNAME")
	atAPIKey := os.Getenv("AT_API_KEY")
	env := os.Getenv("ENV") // e.g., sandbox or production

	client, errClient := clients.NewDefaultHttpClient(&http.Client{})
	if errClient != nil {
		fmt.Println("error client ", errClient.Error())
	}
	//initialise sms
	smsSrv, errSms := services.NewSmsService(atUsername, atAPIKey, env, client)
	if errSms != nil {
		fmt.Println("error client ", errSms.Error())
	}

	//call send
	str, errStr := smsSrv.Send( "", "+254714707147", "Test msg")
	if errStr != nil {
		fmt.Println("error client ", errStr.Error())

	}

	fmt.Printf("sms --> %s", str)

	// Initialize Africa's Talking service
	config.SetSMSService()

	// Example usage of sending an SMS

	// Database connection string
	dsn := "postgres://postgres:mikeserver@localhost:5432/savannah"
	config.InitializeDB(dsn)
	defer config.CloseDB()

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
