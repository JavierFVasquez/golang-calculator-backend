package controllers

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/auth"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/utils"
	"github.com/JavierFVasquez/golang-calculator-backend/v1/record/services"
	"github.com/aws/aws-lambda-go/events"
)

type DeleteRecordController struct {
	service services.RecordServiceIF
}

func NewDeleteRecordController(createService services.RecordServiceIF) DeleteRecordController {
	return DeleteRecordController{
		service: createService,
	}
}

func (controller *DeleteRecordController) DeleteRecordController(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer

	ctxWithValue, err := auth.AuthMiddleware(ctx, request)
	if err != nil {
		return utils.APIError(err, &buf, 401), nil
	}

	recordId := request.PathParameters["id"]

	deletedRecord, operationErr := controller.service.DeleteRecord(*ctxWithValue, &recordId)
	if operationErr != nil {
		return utils.APIError(operationErr, &buf, 400), nil
	}
	body, marshalErr := json.Marshal(deletedRecord)
	if marshalErr != nil {
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
