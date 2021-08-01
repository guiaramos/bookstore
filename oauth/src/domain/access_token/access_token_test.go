package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
  assert.EqualValues(t,24, expirationTime,"expiration time should be 24 hours" )
}

func TestNewAccessToken(t *testing.T) {
	at := NewAccessToken()
  assert.False(t, at == nil, "brand new access token should not be expired")
  assert.EqualValues(t, "", at.AccessToken, "new access token should not have an access token id")
  assert.True(t, at.UserID == 0 ,"new access token should not have an associated user id")
}

func TestNewAccessTokenIsExpired(t *testing.T) {
  at := AccessToken{}
  assert.True(t, at.IsExpired(), "empty access token should be expired by default" )
  at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
  assert.False(t, at.IsExpired(), "access token created three hours from now should NOT be expired")
}
