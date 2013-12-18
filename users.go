package sesame

import (
	"bytes"
	"errors"
	"time"
)

// A User combines Password and Salt into an interface to validate plaintext
// against.
type User struct {
	Email    string   `gorethink:"email"`
	Password Password `gorethink:"password"`
	Salt     Salt     `gorethink:"salt"`

	Created time.Time `gorethink:"created"`
	Updated time.Time `gorethink:"updated"`
}

// SetPassword sets a new hashed password from plaintext
func (u *User) SetPassword(plaintext []byte) error {
	salt, err := NewSalt(20)
	if err != nil {
		return err
	}

	password, err := NewPassword(salt, plaintext)
	if err != nil {
		return err
	}

	u.Salt = salt
	u.Password = password

	return nil
}

// ValidatePassword validates a plaintext password against the stored hash
func (u *User) ValidatePassword(plaintext []byte) (bool, error) {
	hashed, err := NewPassword(u.Salt, plaintext)
	if err != nil {
		return false, err
	}

	return bytes.Equal(hashed, u.Password), nil
}

// ChangePassword is a convenience method to set a new password on a User if
// and only if the original is valid.
func (u *User) ChangePassword(original, updated []byte) error {
	valid, err := u.ValidatePassword(original)
	if !valid {
		return errors.New("invalid original password")
	}
	if err != nil {
		return err
	}

	err = u.SetPassword(updated)
	if err != nil {
		return err
	}

	return nil
}
