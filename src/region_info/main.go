package main

import (
	"encoding/json"
	"fmt"
	"log"
	"region-info/common"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

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
	log.Printf("Handler: %v", request)
	requstBody := Request{}
	err := json.Unmarshal([]byte(request.Body), &requstBody)

	if err != nil {
		log.Println("Error: ", err.Error())
		errorRes := ErrorRes{Message: fmt.Sprintf("Internal Server Error %v", err.Error())}
		return common.ResInternalError(errorRes), err
	}

	resBody := Response{
		MethodRequest: request.HTTPMethod,
		Message:       fmt.Sprintf("Hello %v!", requstBody.Key),
	}
	resByte, err := json.Marshal(resBody)

	if err != nil {
		log.Println("Error2: ", err.Error())
		errorRes := ErrorRes{Message: fmt.Sprintf("Internal Server Error %v", err.Error())}
		return common.ResInternalError(errorRes), nil
	}

	log.Printf("Response: %+v ", string(resByte))
	response := common.ResOk(resByte)
	return response, nil
}

func GETHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Handler: %v", request)
	resBody := Response{
		MethodRequest: request.HTTPMethod,
		Message:       "Hello, GET!",
	}
	resByte, err := json.Marshal(resBody)

	if err != nil {
		log.Println("Error2: ", err.Error())
		errorRes := ErrorRes{Message: fmt.Sprintf("Internal Server Error. %v", err.Error())}
		return common.ResInternalError(errorRes), nil
	}

	log.Printf("Response: %+v ", string(resByte))
	response := common.ResOk(resByte)
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

	return common.ResInternalError(resByte), nil
}

func main() {
	app := NewApp("RegionInfo")
	log.Println("Start Lambda")
	lambda.Start(app.Handler)
}
