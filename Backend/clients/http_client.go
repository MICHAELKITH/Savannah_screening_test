package clients

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// interface
type HttpClientInterface interface{
	PerformHttpCall(req HttpRequest,values url.Values) HttpResponse
}

// HttpRequest models
type HttpRequest struct{
	Method string
	URL string
	Payload interface{}
	Header  []Headers 
}
type Headers struct{
	Key string
	Value string
}
type HttpResponse struct{
	Body []byte
	StatusCode int
	Status string
	Err error
}
type DefaultHttpClient struct{
	client *http.Client
}

func NewDefaultHttpClient(client *http.Client) (*DefaultHttpClient,error){
	if client == nil {
		return nil, errors.New("http client is need")
	}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &DefaultHttpClient{
		client: client,
	},nil
}

func (d *DefaultHttpClient) PerformHttpCall(reqParams HttpRequest,values url.Values) (resp HttpResponse){
	reader := strings.NewReader(values.Encode())

	req,err := http.NewRequest(reqParams.Method,reqParams.URL,reader)
	if err != nil {
		msg := errors.New("error making new request")
		fmt.Println("Error ",err.Error())
		resp.Err = msg
		return resp
	}

	//add headers

	for _,head := range reqParams.Header {
		req.Header.Add(head.Key,head.Value)
	}

	// send request 
	response, errClient := d.client.Do(req)
	if errClient != nil {
		msg := errors.New("error from client call")
		fmt.Println("Error client call ",errClient.Error())
		resp.Err = msg
		return resp
	}
	defer func (Body  io.ReadCloser)  {
		errBody := Body.Close()
		if errBody != nil {
			fmt.Println("response body close error",errBody)
		}
		
	}(response.Body)

	respBody, errRead := io.ReadAll(response.Body)

	if errRead != nil{
		msg := fmt.Errorf("error reading the body values, %v",errRead.Error())
		resp.Err = msg
		resp.StatusCode = http.StatusInternalServerError
		return resp
	}
	resp.Err = nil 
	resp.Body = respBody
	resp.Status = response.Status
	resp.StatusCode = response.StatusCode
	return resp
}