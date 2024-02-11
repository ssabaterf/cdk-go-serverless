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
}
type Response struct {
	Message string            `json:"message"`
	Headers map[string]string `json:"headers"`
}

type ErrorRes struct {
	Message string `json:"message"`
}

func NewApp(id string) *App {
	return &App{id: id}
}

func (app *App) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Handler: %v", request)
	resBody := Response{
		Message: "POOOOOoooonnngggg",
		Headers: request.Headers,
	}
	resByte, err := json.Marshal(resBody)

	if err != nil {
		log.Println("Error: ", err.Error())
		errorRes := ErrorRes{Message: fmt.Sprintf("Internal Server Error. %v", err.Error())}
		return common.ResInternalError(errorRes), nil
	}

	log.Printf("Response: %+v ", string(resByte))
	response := common.ResOk(resByte)
	return response, nil
}

func main() {
	app := NewApp("PingLambda")
	log.Println("Start Lambda")
	lambda.Start(app.Handler)
}
