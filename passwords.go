package sesame

import "code.google.com/p/go.crypto/bcrypt"

const Cost = 10

type Password []byte

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

func NewPassword(s Salt, plaintext []byte) (password Password, err error) {
	password, err = bcrypt.GenerateFromPassword(append(s, plaintext...), Cost)

	defer clear(plaintext)
	return
}
