package sesame

import (
	"errors"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestValidateEmailInvalid(t *testing.T) {
	err := ValidateEmail("bob")

	if assert.NotNil(t, err) {
		assert.Equal(t, err, errors.New("email address is not valid: no @"))
	}
}

func TestValidateEmailValid(t *testing.T) {
	err := ValidateEmail("a@b.com")

	assert.Nil(t, err)
}
