package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/auth"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/services"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/utils"
	"github.com/aws/aws-lambda-go/events"
)

type DivisionController struct {
	service services.BasicOperationServiceIF
}

func NewDivisionController(createService services.BasicOperationServiceIF) DivisionController {
	return DivisionController{
		service: createService,
	}
}

func (controller *DivisionController) DivisionController(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctxWithValue, err := auth.AuthMiddleware(ctx, request)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 401}, nil
	}

	var buf bytes.Buffer

	var operation models.Operation
	if err := json.Unmarshal([]byte(request.Body), &operation); err != nil {
		return utils.APIError(&err, &buf, 400), nil
	}

	operation.Operation = models.DIVISION
	if operation.Operand2 == 0 {
		divisionByZeroErr := errors.New("DIVISION_BY_ZERO_ERROR")
		return utils.APIError(&divisionByZeroErr, &buf, 400), nil
	}

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
