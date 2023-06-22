package controllers

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/auth"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/utils"
	"github.com/JavierFVasquez/truenorth-calculator-backend/v1/randomString/services"
	"github.com/aws/aws-lambda-go/events"
)

type RandomStringController struct {
	service services.RandomStringServiceIF
}

func NewRandomStringController(createService services.RandomStringServiceIF) RandomStringController {
	return RandomStringController{
		service: createService,
	}
}

func (controller *RandomStringController) RandomStringController(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctxWithValue, err := auth.AuthMiddleware(ctx, request)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 401}, nil
	}

	var buf bytes.Buffer

	var operation models.Operation
	if err := json.Unmarshal([]byte(request.Body), &operation); err != nil {
		return utils.APIError(&err, &buf, 400), nil
	}

	operation.Operation = models.RANDOM_STRING

	record, operationErr := controller.service.RandomString(*ctxWithValue, operation)
	if operationErr != nil {
		errorResponse := map[string]interface{}{
			"error": (*operationErr).Error(),
		}
		body, marshalErr := json.Marshal(errorResponse)
		if marshalErr != nil {
			return utils.APIError(&marshalErr, &buf, 400), nil
		}
		json.HTMLEscape(&buf, body)
		return events.APIGatewayProxyResponse{
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
		return events.APIGatewayProxyResponse{StatusCode: 500}, marshalErr
	}
	json.HTMLEscape(&buf, body)

	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}