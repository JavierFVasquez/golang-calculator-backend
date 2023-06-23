package utils

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func APIError(paramError *error, buf *bytes.Buffer, statusCode int) events.APIGatewayProxyResponse {

	errorResponse := map[string]interface{}{
		"errorMessage": (*paramError).Error(),
	}
	body, marshalErr := json.Marshal(errorResponse)
	if marshalErr != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}
	json.HTMLEscape(buf, body)
	return events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            (*buf).String(),
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Headers": "Content-Type, Authorization",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
		},
	}
}
