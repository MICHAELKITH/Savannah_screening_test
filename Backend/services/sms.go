package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/savannah/sms/clients"
	"github.com/savannah/sms/models"
)

const (
	Sandbox = "sandbox"
	Prod    = "production"
	baseUrl = "https://api.sandbox.africastalking.com"
)

type SmsInterface interface{
	Send(from,to, message string) (*models.SendMessageResponse, error)
}

// Service is a model
type Service struct {
	Username string
	APIKey   string
	Env      string
	Client   clients.HttpClientInterface
}

// NewService returns a new service
func NewSmsService(username, apiKey, env string, client clients.HttpClientInterface) (*Service, error) {
	if username == "" {
		return nil, errors.New("Username is required")
	}
	if apiKey == "" {
		return nil, errors.New("apiKey is required")
	}
	if env == "" {
		return nil, errors.New("env is required")
	}
	if client == nil {
		return nil, errors.New("client is required and shouldn't be nil")
	}
	return &Service{username, apiKey, env, client}, nil
}

// Send - POST
func (service *Service) Send(from,to, message string) (*models.SendMessageResponse, error) {
	values := url.Values{}
	values.Set("username", service.Username)
	values.Set("to", to)
	values.Set("message", message)
	if from != "" {
		// set from = "" to avoid this
		values.Set("from", from)
	}

	urlParse, errUrl := constructUrlWithParams(baseUrl, "/version1/messaging", nil)
	if errUrl != nil {
		fmt.Println("error url construct")
		return nil,errUrl
	}

	fmt.Println("constructed URL", urlParse)
	header := []clients.Headers{
		{Key: "Content-Type", Value: "application/x-www-form-urlencoded"},
		{Key: "apikey", Value: service.APIKey},
		{Key: "Accept", Value: "application/json"},
	}
	args := clients.HttpRequest{
		Method:  http.MethodPost,
		URL:     urlParse,
		Header:  header,
		Payload: nil,
	}

	resp := service.Client.PerformHttpCall(args,values)

	if resp.Err != nil {
		fmt.Println("Error", resp.Err)
		return nil,resp.Err
	}

	fmt.Println("Body results", string(resp.Body))

	
	var smsMessageResponse models.SendMessageResponse
	if err := json.Unmarshal(resp.Body,&smsMessageResponse); err != nil {
		return nil, errors.New("unable to parse sms response")
	}
	fmt.Println("results", smsMessageResponse)

	return &smsMessageResponse, nil
}
func constructUrlWithParams(baseUrl, endPoint string, params map[string]string) (string, error) {
	u, err := url.Parse(baseUrl + endPoint)
	if err != nil {
		return "", err
	}
	//create query params
	query := url.Values{}

	for key, value := range params {
		query.Add(key, value)
	}
	//append the query params into url
	u.RawQuery = query.Encode()
	return u.String(), nil

}
