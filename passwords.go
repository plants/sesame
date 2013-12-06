package sesame

import "code.google.com/p/go.crypto/scrypt"

const (
	scrypt_N      = 16384
	scrypt_r      = 8
	scrypt_p      = 1
	scrypt_keyLen = 64
)

type Password []byte

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

func NewPassword(salt Salt, plaintext []byte) (password Password, err error) {
	password, err = scrypt.Key(plaintext, salt, scrypt_N, scrypt_r, scrypt_p, scrypt_keyLen)

	defer clear(plaintext)
	return
}
