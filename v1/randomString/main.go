package main

import (
	"context"
	"os"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/clients"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/constants"
	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/repositories"
	"github.com/JavierFVasquez/truenorth-calculator-backend/v1/randomString/controllers"
	"github.com/JavierFVasquez/truenorth-calculator-backend/v1/randomString/repositories"
	"github.com/JavierFVasquez/truenorth-calculator-backend/v1/randomString/services"
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
	randomStringRepository := random_string_repositories.NewRandomStringRepository(&logger)
	operationRepository := repositories.NewOperationRepository(mongoClient, string(constants.OperationCollection), &logger)
	randomStringService := services.NewRandomStringService(recordRepository, operationRepository, randomStringRepository, &logger)

	randomStringController := controllers.NewRandomStringController(randomStringService)

	lambda.Start(randomStringController.RandomStringController)
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
