package sms_test

import (
	"testing"
	"github.com/MICHAELKITH/service/config"
	"github.com/stretchr/testify/assert"
)

func TestSendSMS(t *testing.T) {
	// Set up mock phone number and message
	phoneNumber := "+254714707147"
	message := "Hello from test!"

	// Assuming config.SetSMSService() initializes the SMS service
	config.SetSMSService()
	
	err := config.SendSMS(phoneNumber, message)
	assert.NoError(t, err, "expected no error in sending SMS")

	
}
