package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/auth"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/utils"
	"github.com/JavierFVasquez/golang-calculator-backend/v1/record/services"
	"github.com/aws/aws-lambda-go/events"
)

type RecordController struct {
	service services.RecordServiceIF
}

func NewRecordController(createService services.RecordServiceIF) RecordController {
	return RecordController{
		service: createService,
	}
}

func (controller *RecordController) RecordController(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer

	ctxWithValue, err := auth.AuthMiddleware(ctx, request)
	if err != nil {
		return utils.APIError(err, &buf, 401), nil
	}

	options := utils.ParsePaginationParams(request.QueryStringParameters)

	recordListPaginated, operationErr := controller.service.GetRecordList(*ctxWithValue, options)
	if operationErr != nil {
		return utils.APIError(operationErr, &buf, 400), nil
	}
	fmt.Println("recordListPaginated", recordListPaginated)
	body, marshalErr := json.Marshal(recordListPaginated)
	if marshalErr != nil {
		fmt.Println("marshalErr", marshalErr)
		return utils.APIError(&marshalErr, &buf, 400), nil
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
