package clients

import (
	"context"
	"errors"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/rs/zerolog"
	"google.golang.org/api/option"
)

var (
	ErrFirebaseAuth = errors.New("Firebase authentication error")
)

type FirebaseClient struct {
	auth   *auth.Client
	logger *zerolog.Logger
}

func NewFirebaseClient(ctx context.Context, logger *zerolog.Logger) *FirebaseClient {
	log := logger.With().Str("resource", "firebase_client").Logger()

	serviceAccount, fileExists := os.LookupEnv("GOOGLE_SERVICE_ACCOUNT")
	if !fileExists {
		log.Panic().Msg("Firebase credentials not found")
	}

	opt := option.WithCredentialsFile(serviceAccount)
	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Panic().Err(err).Msg("Unable to initialize firebase client")
	}

	auth, err := fbApp.Auth(ctx)
	if err != nil {
		log.Panic().Err(err).Msg("Firebase Auth client couldn't be initialized")
	}

	log.Info().Msg("Successfully connected to Firebase!")

	return &FirebaseClient{
		auth:   auth,
		logger: &log,
	}
}

func (client *FirebaseClient) Auth(ctx context.Context, token string) (*auth.Token, error) {
	user, err := client.auth.VerifyIDToken(ctx, token)
	if err != nil {
		client.logger.Error().Err(err).Msg("Firebase authentication error")
		return nil, ErrFirebaseAuth
	}

	return user, nil
}
