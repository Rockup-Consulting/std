package hashx

import (
	"golang.org/x/crypto/bcrypt"
)

func BcryptNew(value string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}

	return string(hashedBytes)
}

func BcryptNewLowCost(value string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}

	return string(hashedBytes)
}

func BcryptCompare(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password))

	switch err {
	case nil:
		return true, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return false, nil
	default:
		return false, err
	}
}
