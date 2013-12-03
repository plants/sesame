package sesame

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestNewSalt(t *testing.T) {
	salt, err := NewSalt(20)

	assert.Nil(t, err)
	assert.Equal(t, len(salt), 20)
}
