package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/JavierFVasquez/golang-calculator-backend/libs/models"
	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func AuthMiddleware(ctx context.Context, request events.APIGatewayProxyRequest) (*context.Context, *error) {
	authToken := strings.Replace(request.Headers["authorization"], "Bearer ", "", 1)
	fmt.Println("Headers=", request.Headers)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	input := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(authToken),
	}

	result, err := cognitoClient.GetUser(input)
	if err != nil {
		fmt.Println("Error al obtener el usuario:", err)
		err := errors.New("INVALID_TOKEN")
		return nil, &err
	}
	var userId string
	var userName string
	var userEmail string

	for _, att := range result.UserAttributes {
		switch *att.Name {
		case "email":
			userEmail = *att.Value
		case "name":
			userName = *att.Value
		case "sub":
			userId = *att.Value
		}
	}

	authUser := models.User{
		ID:    userId,
		Name:  userName,
		Email: userEmail,
	}
	ctxWithValue := context.WithValue(ctx, "user", authUser)

	return &ctxWithValue, nil
}
