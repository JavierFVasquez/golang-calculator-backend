package random_string_repositories

import (
	"context"
)

//go:generate mockery --name RandomStringRepositoryIF --structname MockRandomStringRepositoryIF --filename record_repository_mock.go
type RandomStringRepositoryIF interface {
	GetRandomString(ctx context.Context) (*string, error)
}
