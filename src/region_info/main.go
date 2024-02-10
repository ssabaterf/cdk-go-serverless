package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var headersError = map[string]string{"Content-Type": "application/json"}
var headersResponse = map[string]string{
	"Content-Type":                     "application/json",
	"Access-Control-Allow-Origin":      "*",
	"Access-Control-Allow-Methods":     "GET, POST, PUT, DELETE, OPTIONS, HEAD",
	"Access-Control-Allow-Headers":     "Content-Type, Authorization, X-Amz-Date, X-Api-Key, X-Amz-Security-Token",
	"Access-Control-Allow-Credentials": "true",
}

type App struct {
	id string
}

type Request struct {
	Key string `json:"key"`
}
type Response struct {
	MethodRequest string `json:"methodRequest"`
	Message       string `json:"message"`
}

type ErrorRes struct {
	Message string `json:"message"`
}

func NewApp(id string) *App {
	log.Println("NewApp")
	return &App{id: id}
}

func POSTHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("POSTHandler")
	requstBody := Request{}
	err := json.Unmarshal([]byte(request.Body), &requstBody)

	if err != nil {
		log.Println("Error: ", err.Error())
		text := fmt.Sprintf("Internal Server Error. %v", err.Error())
		errorRes := ErrorRes{Message: text}
		resByte, _ := json.Marshal(errorRes)
		return events.APIGatewayProxyResponse{
			Body:       string(resByte),
			StatusCode: http.StatusInternalServerError,
			Headers:    headersError,
		}, errors.New(text)
	}

	resBody := Response{
		MethodRequest: request.HTTPMethod,
		Message:       "Hello, " + requstBody.Key + "!",
	}
	resByte, err := json.Marshal(resBody)

	if err != nil {
		log.Println("Error2: ", err.Error())
		text := fmt.Sprintf("Internal Server Error. %v", err.Error())
		errorRes := ErrorRes{Message: text}
		resByte, _ := json.Marshal(errorRes)
		return events.APIGatewayProxyResponse{
			Body:       string(resByte),
			StatusCode: http.StatusInternalServerError,
			Headers:    headersError,
		}, nil
	}

	log.Printf("Response: %+v ", string(resByte))
	response := events.APIGatewayProxyResponse{
		Body:       string(resByte),
		StatusCode: http.StatusOK,
		Headers:    headersResponse,
	}
	return response, nil
}

func GETHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("GETHandler")
	resBody := Response{
		MethodRequest: request.HTTPMethod,
		Message:       "Hello, GET!",
	}
	resByte, err := json.Marshal(resBody)

	if err != nil {
		log.Println("Error2: ", err.Error())
		text := fmt.Sprintf("Internal Server Error. %v", err.Error())
		errorRes := ErrorRes{Message: text}
		resByte, _ := json.Marshal(errorRes)
		return events.APIGatewayProxyResponse{
			Body:       string(resByte),
			StatusCode: http.StatusInternalServerError,
			Headers:    headersError,
		}, nil
	}

	log.Printf("Response: %+v ", string(resByte))
	response := events.APIGatewayProxyResponse{
		Body:       string(resByte),
		StatusCode: http.StatusOK,
		Headers:    headersResponse,
	}
	return response, nil
}

func (app *App) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Handler: %v", app)

	if request.HTTPMethod == "POST" {
		return POSTHandler(request)
	} else if request.HTTPMethod == "GET" {
		return GETHandler(request)
	}

	errorRes := ErrorRes{Message: "Not Found " + request.HTTPMethod}
	resByte, _ := json.Marshal(errorRes)

	return events.APIGatewayProxyResponse{
		Body:       string(resByte),
		StatusCode: http.StatusInternalServerError,
		Headers:    headersError,
	}, nil
}

func main() {
	log.Println("Start RegionInfo")
	app := NewApp("RegionInfo")
	log.Println("Start Lambda")
	lambda.Start(app.Handler)
}
