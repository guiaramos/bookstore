package access_token

import (
	"time"
)

const (
	expirationTime = 24
)

// AccessToken struct represents the access token informations
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// NewAccessToken function creates a new AccessToken
func NewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired method checks if the access token is expired
func (a AccessToken) IsExpired() bool {
	return time.Unix(a.Expires, 0).Before(time.Now().UTC())
}
