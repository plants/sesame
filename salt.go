package auth

import (
	"crypto/rand"
	"io"
)

type Salt []byte

func NewSalt(length int) (salt Salt, err error) {
	salt = make(Salt, length)
	_, err = io.ReadFull(rand.Reader, salt)
	return
}
