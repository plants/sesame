package sesame

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestNewPassword(t *testing.T) {
	pass, err := NewPassword(Salt("salt"), []byte("password"))

	assert.Nil(t, err)
	assert.Equal(t, len(pass), 64)
}

func TestRepeatability(t *testing.T) {
	salt := Salt("salt")

	a, err := NewPassword(salt, []byte("password"))
	assert.Nil(t, err)
	b, err := NewPassword(salt, []byte("password"))
	assert.Nil(t, err)

	assert.Equal(t, a, b)
}

func TestClear(t *testing.T) {
	value := []byte("test")

	clear(value)

	assert.Equal(t, value, []byte{0, 0, 0, 0})
}
