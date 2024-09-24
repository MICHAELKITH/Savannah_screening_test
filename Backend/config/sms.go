package config

import (
    "log"
    "os"

    "github.com/AndroidStudyOpenSource/africastalking-go/sms"
   
)

var smsService sms.Service

// SetSMSService initializes the SMS service with credentials from environment variables
func SetSMSService() {
    // Get credentials from environment variables
    atUsername := os.Getenv("AT_USERNAME")
    atAPIKey := os.Getenv("AT_API_KEY")
    env := os.Getenv("ENV") // e.g., sandbox or production

    if atUsername == "" || atAPIKey == "" {
        log.Fatal("AT_USERNAME and AT_API_KEY must be set")
    }

    // Initialize Africa's Talking service
    smsService = sms.NewService(atUsername, atAPIKey, env)
}

// GetSMSService returns the initialized SMS service
func GetSMSService() sms.Service {
    return smsService
}
