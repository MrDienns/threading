package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

type BCrypt struct {
}

func NewBCrypt() *BCrypt {
	return &BCrypt{}
}

func (b *BCrypt) Compare(hashed, original string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(original)) == nil
}
