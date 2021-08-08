package db

import (
	"github.com/gocql/gocql"
	"github.com/guiaramos/bookstore/oauth/src/clients/cassandra"
	"github.com/guiaramos/bookstore/oauth/src/domain/access_token"
	"github.com/guiaramos/bookstore/oauth/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

// DBRepository interface represents the DB repository
type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

// NewDBRepository function creates a new DB Repository
func NewDBRepository() DBRepository {
	return &dbRepository{}
}

// Create inserts a new access token to the database
func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInterServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken, at.AccessToken, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.NewInterServerError(err.Error())
	}

	return nil
}

// GetByID method gets an AccessToken by ID
func (r dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInterServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id.")
		}
		return nil, errors.NewInterServerError(err.Error())
	}

	return &result, nil
}

// UpdateExpirationTime updates the expiration time of an access token in the database
func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInterServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInterServerError(err.Error())
	}

	return nil
}
