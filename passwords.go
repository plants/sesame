package sesame

import "code.google.com/p/go.crypto/scrypt"

const (
	scryptN      = 16384
	scryptr      = 8
	scryptp      = 1
	scryptkeyLen = 64
)

// Password stores an encrypted password
type Password []byte

// zero out a value in memory
func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

// NewPassword creates a new password, using the constants defined in this package.
func NewPassword(salt Salt, plaintext []byte) (password Password, err error) {
	password, err = scrypt.Key(plaintext, salt, scryptN, scryptr, scryptp, scryptkeyLen)

	defer clear(plaintext)
	return
}
