package sesame

import (
	"crypto/rand"
	"io"
)

// A Salt represents a cryptographically random salt value
type Salt []byte

// NewSalt generates cryptographically random salts
func NewSalt(length int) (salt Salt, err error) {
	salt = make(Salt, length)
	_, err = io.ReadFull(rand.Reader, salt)
	return
}
