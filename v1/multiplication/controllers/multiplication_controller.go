package controllers

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/auth"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/services"
	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

type MultiplicationController struct {
	service services.BasicOperationServiceIF
}

func NewMultiplicationController(createService services.BasicOperationServiceIF) MultiplicationController {
	return MultiplicationController{
		service: createService,
	}
}

func (controller *MultiplicationController) MultiplicationController(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	ctxWithValue, err := auth.AuthMiddleware(ctx, request)
	if err != nil {
		return Response{StatusCode: 401}, nil
	}

	var buf bytes.Buffer

	var operation models.Operation
	if err := json.Unmarshal([]byte(request.Body), &operation); err != nil {
		return Response{StatusCode: 400}, err
	}

	operation.Operation = models.MULTIPLICATION

	record, operationErr := controller.service.BasicOperation(*ctxWithValue, operation)
	if operationErr != nil {
		errorResponse := map[string]interface{}{
			"error": (*operationErr).Error(),
		}
		body, marshalErr := json.Marshal(errorResponse)
		if marshalErr != nil {
			return Response{StatusCode: 400}, marshalErr
		}
		json.HTMLEscape(&buf, body)
		return Response{
			StatusCode:      400,
			IsBase64Encoded: false,
			Body:            buf.String(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}
	body, marshalErr := json.Marshal(record)
	if marshalErr != nil {
		return Response{StatusCode: 500}, marshalErr
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}
