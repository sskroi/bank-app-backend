package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

// Password Hasher
type PasswdHasher interface {
	Hash(password string) (string, error)
	Check(hashedPassowrd, password string) bool
}

type BcryptHasher struct {
	cost int
}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{cost: bcrypt.DefaultCost}
}

// max len 72 bytes
func (h *BcryptHasher) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (h *BcryptHasher) Check(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
