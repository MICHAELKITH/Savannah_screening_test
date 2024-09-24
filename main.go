package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
		log.Fatalf("Error creating HTTP client: %v", errClient)
	}

	// Initialize SMS service
	smsSrv, errSms := services.NewSmsService(atUsername, atAPIKey, env, client)
	if errSms != nil {
		log.Fatalf("Error initializing SMS service: %v", errSms)
	}

	// Call send (example number, change as necessary)
	str, errStr := smsSrv.Send("", "+254714707147", "Test msg")
	if errStr != nil {
		log.Fatalf("Error sending SMS: %v", errStr)
	}
	fmt.Printf("SMS sent: %s\n", str)

	// Initialize Africa's Talking service
	config.SetSMSService()

	// Get DB credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Build the connection string
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Database connection
	config.InitializeDB(dsn)
	defer config.CloseDB()

	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS with customized settings
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Replace "*" with your frontend URL for production
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":8000"))
}
