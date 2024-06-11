package tests

import (
	"testing"

	"github.com/saddam-satria/posq-be/commons"
	"github.com/stretchr/testify/assert"
)

func TestBcrypt(t *testing.T) {
	password := "test123"

	hashedPassword, err := commons.HashedBcrypt(password)

	assert.NoError(t, err)

	isMatch := commons.VerifyHashed(password, hashedPassword)

	assert.Equal(t, true, isMatch)
}