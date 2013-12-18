package sesame

import (
	"errors"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

func TestSetPassword(t *testing.T) {
	u := new(User)
	u.SetPassword([]byte("valid"))

	assert.Equal(t, len(u.Salt), 20)
	assert.Equal(t, len(u.Password), 64)
}

func TestValidatePassword(t *testing.T) {
	u := new(User)
	u.SetPassword([]byte("valid"))

	// a password with a byte string equal to what we passed in before should
	// pass validation
	good, err := u.ValidatePassword([]byte("valid"))
	assert.True(t, good)
	assert.Nil(t, err)

	// a password with a byte string not equal to what we passed in before
	// should not pass validation
	bad, err := u.ValidatePassword([]byte("invalid"))
	assert.False(t, bad)
	assert.Nil(t, err)
}

func TestChangePassword(t *testing.T) {
	u := new(User)
	u.SetPassword([]byte("valid"))

	// changing the password fails for an invalid password
	err := u.ChangePassword([]byte("invalid"), []byte("newpass"))
	assert.Equal(t, err, errors.New("invalid original password"))

	// changing the password works if you pass in a valid password
	err = u.ChangePassword([]byte("valid"), []byte("newpass"))
	assert.Nil(t, err)
}

func TestNewUser(t *testing.T) {
	u := NewUser("test@example.com", "password")

	assert.Equal(t, "test@example.com", u.Email)
	assert.NotEqual(t, new(Password), u.Password)
	assert.NotEqual(t, new(Salt), u.Salt)

	assert.NotEqual(t, time.Time{}, u.Created)
	assert.NotEqual(t, time.Time{}, u.Updated)
	assert.Equal(t, u.Created, u.Updated)
}
