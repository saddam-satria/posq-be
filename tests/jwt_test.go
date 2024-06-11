package tests

import (
	"testing"
	"time"

	"github.com/saddam-satria/posq-be/commons"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	id := "123"
	email := "test@gmail.com"
	expired := time.Now().Add(time.Minute * 60 * 24).Unix()
	token, err := commons.GenerateToken(id, email, expired, "inisecretkey")

	assert.NoError(t, err, "error generate token")

	claims, err := commons.VerifyToken(token)

	assert.NoError(t, err, "error verifiy token")

	assert.Equal(t, id, claims["id"], token)
}