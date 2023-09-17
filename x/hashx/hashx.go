package hashx

import (
	"crypto/md5"
	"fmt"
	"io"

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

func MD5(r io.Reader) (string, error) {
	h := md5.New()

	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
