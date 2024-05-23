package controllers

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/auth"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/models"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/services"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/utils"
	"github.com/aws/aws-lambda-go/events"
)

type MultiplicationController struct {
	service services.BasicOperationServiceIF
}

func NewMultiplicationController(createService services.BasicOperationServiceIF) MultiplicationController {
	return MultiplicationController{
		service: createService,
	}
}

func (controller *MultiplicationController) MultiplicationController(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctxWithValue, err := auth.AuthMiddleware(ctx, request)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 401}, nil
	}

	var buf bytes.Buffer

	var operation models.Operation
	if err := json.Unmarshal([]byte(request.Body), &operation); err != nil {
		return utils.APIError(&err, &buf, 400), nil
	}

	operation.Operation = models.MULTIPLICATION

	record, operationErr := controller.service.BasicOperation(*ctxWithValue, operation)
	if operationErr != nil {
		return utils.APIError(operationErr, &buf, 412), nil
	}
	body, marshalErr := json.Marshal(record)
	if marshalErr != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, marshalErr
	}
	json.HTMLEscape(&buf, body)

	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
	return resp, nil
}
