package sesame

import (
	"errors"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestInMemoryStore(t *testing.T) {
	store, _ := NewInMemoryStore()
	user := User{Email: "test@example.com", Password: []byte("test"), Salt: []byte("salt")}

	// GetByEmail without a user
	u, err := store.GetByEmail("test@example.com")
	assert.Nil(t, u)
	assert.Equal(t, err, errors.New(NoSuchUserError))

	// Save
	err = store.Save(&user)
	assert.Nil(t, err)

	// GetByEmail with a user
	u, err = store.GetByEmail("test@example.com")
	assert.Equal(t, u, &user)
}
