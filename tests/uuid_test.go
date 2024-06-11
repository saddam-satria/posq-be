package tests

import (
	"testing"

	"github.com/saddam-satria/posq-be/commons"
	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {

	generatedUUID := commons.GenerateUUID()

	isValid := commons.IsUUID(generatedUUID)

	assert.Equal(t, true, isValid)
}