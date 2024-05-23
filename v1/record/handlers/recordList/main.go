package main

import (
	"context"
	"os"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/clients"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/constants"
	"github.com/JavierFVasquez/golang-calculator-backend/libs/repositories"
	"github.com/JavierFVasquez/golang-calculator-backend/v1/record/controllers"
	"github.com/JavierFVasquez/golang-calculator-backend/v1/record/services"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	ctx := context.Background()

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Env
	mongoUri := getEnv("MONGO_URI", "")
	dbName := getEnv("DB_NAME", "")

	// Clients
	mongoClient, mongoClose := clients.NewMongoClient(ctx, mongoUri, dbName, &logger)
	defer mongoClose()

	recordRepository := repositories.NewRecordRepository(mongoClient, string(constants.RecordCollection), &logger)
	recordService := services.NewRecordService(recordRepository, &logger)

	recordController := controllers.NewRecordController(recordService)

	lambda.Start(recordController.RecordController)
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
