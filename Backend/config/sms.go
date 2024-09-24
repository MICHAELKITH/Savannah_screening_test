package config

import (
    "log"
    "os"

    "github.com/AndroidStudyOpenSource/africastalking-go/sms"
   
)

var smsService sms.Service

// credentials from environment variables
func SetSMSService() {
    // environment variables
    atUsername := os.Getenv("AT_USERNAME")
    atAPIKey := os.Getenv("AT_API_KEY")
    env := os.Getenv("ENV")

    if atUsername == "" || atAPIKey == "" {
        log.Fatal("AT_USERNAME and AT_API_KEY must be set")
    }

    smsService = sms.NewService(atUsername, atAPIKey, env)
}


func GetSMSService() sms.Service {
    return smsService
}
