package common

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

var HeadersError = map[string]string{"Content-Type": "application/json"}
var HeadersResponse = map[string]string{
	"Content-Type":                     "application/json",
	"Access-Control-Allow-Origin":      "*",
	"Access-Control-Allow-Methods":     "GET, POST, PUT, DELETE, OPTIONS, HEAD",
	"Access-Control-Allow-Headers":     "Content-Type, Authorization, X-Amz-Date, X-Api-Key, X-Amz-Security-Token",
	"Access-Control-Allow-Credentials": "true",
}

func ResInternalError(resBody interface{}) events.APIGatewayProxyResponse {
	resByte, _ := json.Marshal(resBody)
	return events.APIGatewayProxyResponse{
		Body:       string(resByte),
		StatusCode: http.StatusInternalServerError,
		Headers:    HeadersError,
	}
}

func ResOk(resBody []byte) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       string(resBody),
		StatusCode: http.StatusOK,
		Headers:    HeadersResponse,
	}
}
