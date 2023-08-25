package password

import "golang.org/x/crypto/bcrypt"

func Hash(passwordToHash []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(passwordToHash, bcrypt.MinCost)
}

func Compare(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
