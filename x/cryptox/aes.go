// Package encrypt is a small wrapper package that simplifies usage of encryption tooling.
// Specifically, we wrap the Advanced Encryption Standard cipher in a service called
// encrypt.Service.
//
// To learn more about AES in Go, refer to: https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Service wraps an aes cipher for simplified usage.
type Service struct {
	gcm cipher.AEAD
}

// NewService creates and returns a new encrypt.Service, or an error. To create a service a secret
// of exactly 16, 24 or 32 bytes has to be provided.
//
// If the provided secret is not of the right length, ErrInvalidSecret is returned. This can be
// verified using the encrypt.IsInvalidSecretErr helper.
func NewService(secret string) (Service, error) {
	secretLen := len(secret)

	if !(secretLen == 16 || secretLen == 24 || secretLen == 32) {
		return Service{}, errInvalidSecret(secretLen)
	}

	c, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return Service{}, err
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return Service{}, err
	}

	return Service{
		gcm: gcm,
	}, nil
}

// Encrypt is a method on Service that encrypts a value. See 'Decrypt' to reverse the encryption.
func (s Service) Encrypt(val string) string {

	text := []byte(val)

	nonce := make([]byte, s.gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	encryptedValue := s.gcm.Seal(nonce, nonce, text, nil)

	return string(encryptedValue)
}

// Decrypt is a method on Service that reverses the encryption process. Either the unencrypted value
// is returned, or an error.
//
// If the error is due to an invalid value, the error will be ErrInvalidVal, this can be verified
// using the encrypt.IsInvalidValErr helper
func (s Service) Decrypt(val string) (string, error) {

	text := []byte(val)

	nonceSize := s.gcm.NonceSize()
	if len(text) < nonceSize {
		return "", ErrInvalidVal
	}

	nonce, ciphertext := text[:nonceSize], text[nonceSize:]
	plaintext, err := s.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", ErrInvalidVal
	}

	return string(plaintext), nil
}
