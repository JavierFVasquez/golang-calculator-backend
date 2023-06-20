package controllers

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/services"
	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

type AdditionController struct {
	service services.BasicOperationServiceIF
}

func NewAdditionController(createService services.BasicOperationServiceIF) AdditionController {
	return AdditionController{
		service: createService,
	}
}

type key string


func (controller *AdditionController) AdditionController(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	dummyUser := models.User{
		ID:"ABC123",
		Name: "Javier",
		Email: "javier@example.com",
	}
	ctxWithValue := context.WithValue(ctx, key("user"), dummyUser)

	var operation models.Operation
	if err := json.Unmarshal([]byte(request.Body), &operation); err != nil {
			return Response{StatusCode: 400}, err
	}

	operation.Operation = models.ADDITION

	record, operationErr := controller.service.BasicOperation(ctxWithValue, operation)
	if operationErr != nil {
			return Response{StatusCode: 400}, *operationErr
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
			"Content-Type":           "application/json",
		},
	}
	return resp, nil
}

