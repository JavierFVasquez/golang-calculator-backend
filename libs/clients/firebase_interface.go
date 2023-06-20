package clients

import (
	"context"

	"firebase.google.com/go/auth"
)

//go:generate mockery --name FirebaseClientIF --structname MockFirebaseClientIF --filename firebase_client_mock.go
type FirebaseClientIF interface {
	Auth(ctx context.Context, token string) (*auth.Token, error)
}
