package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash for generate hash from plain
func Hash(plain string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), 14)
	return string(b), err
}

// Compare for check hash and plain is same, with return error
func Compare(hash string, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}

// IsValid for check hash and plain is same, with return bool
func IsValid(hash string, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}
