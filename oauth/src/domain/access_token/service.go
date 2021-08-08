package access_token

import (
	"strings"

	"github.com/guiaramos/bookstore/oauth/utils/errors"
)

// Service interface represents the AccessToken service
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(token AccessToken) *errors.RestErr
	UpdateExpirationTime(token AccessToken) *errors.RestErr
}

// Repository interface represents the AccessToken repository
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(token AccessToken) *errors.RestErr
	UpdateExpirationTime(token AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

// NewService function creates a new Service
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// GetByID method get Access Token by its ID
func (s service) GetByID(id string) (*AccessToken, *errors.RestErr) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.NewBadRequestError("invalid access token")
	}

	return s.repository.GetByID(id)
}

// Create method creates new access token
func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

// UpdateExpirationTime method updates expiration time of access token
func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
