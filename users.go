package sesame

import (
	"bytes"
	"errors"

	"github.com/eaigner/hood"
)

type User struct {
	id       hood.Id
	Email    string   `sql:"size:250,notnull"`
	Password Password `sql:"size:64,notnull"`
	Salt     Salt     `sql:"size:40,notnull"`

	Created hood.Created
	Updated hood.Updated
}

// Set a password.
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

// Validate a password, using the stored salt for both comparisons
func (u *User) ValidatePassword(plaintext []byte) (bool, error) {
	hashed, err := NewPassword(u.Salt, plaintext)
	if err != nil {
		return false, err
	}

	return bytes.Equal(hashed, u.Password), nil
}

// Set a password, first validating that the old password was valid.
func (u *User) ChangePassword(original, updated []byte) error {
	valid, err := u.ValidatePassword(original)
	if !valid {
		return errors.New("Invalid original password.")
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
