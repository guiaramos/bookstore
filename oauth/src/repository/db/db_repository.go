package db

import (
	"github.com/guiaramos/bookstore/oauth/src/clients/cassandra"
	"github.com/guiaramos/bookstore/oauth/src/domain/access_token"
	"github.com/guiaramos/bookstore/oauth/utils/errors"
)

// DBRepository interface representes the DB repository
type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

// NewDBRepository function creates a new DB Repository
func NewDBRepository() DBRepository {
	return &dbRepository{}
}

// GetByID method gets an AccessToken by ID
func (r *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors.NewInterServerError("database connection not implemented yet")
}
