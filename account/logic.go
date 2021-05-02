package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

// NewService - to implement Service interface
func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

//CreateUser - service struct implementing Service interface methods
func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	err := s.repository.CreateUser(ctx, user)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return "", err
	}
	logger.Log("Created user", id)

	return "Success", nil
}

// GetUser - service struct implementing Service interface.
func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")
	email, err := s.repository.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err")
		return "", err
	}

	logger.Log("Get user", id)
	return email, nil
}
