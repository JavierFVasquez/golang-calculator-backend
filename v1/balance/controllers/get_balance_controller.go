package controllers

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/auth"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/utils"
	"github.com/JavierFVasquez/truenorth-calculator-backend/v1/balance/services"
	"github.com/aws/aws-lambda-go/events"
)

type GetBalanceController struct {
	service services.NewBalanceServiceIF
}

func NewGetBalanceController(createService services.NewBalanceServiceIF) GetBalanceController {
	return GetBalanceController{
		service: createService,
	}
}

func (controller *GetBalanceController) GetBalanceController(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctxWithValue, err := auth.AuthMiddleware(ctx, request)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 401}, nil
	}

	var buf bytes.Buffer

	userWithBalance, operationErr := controller.service.GetBalance(*ctxWithValue)
	if operationErr != nil {
		return utils.APIError(operationErr, &buf, 412), nil
	}
	body, marshalErr := json.Marshal(userWithBalance)
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
