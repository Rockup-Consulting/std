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
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

// AES
// ====================================================================

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

// RSA
// ====================================================================

func RsaEncrypt(secret string) (string, string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err.Error())
	}

	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &privateKey.PublicKey, []byte(secret), label)
	if err != nil {
		panic(err.Error())
	}

	return base64.StdEncoding.EncodeToString(ciphertext), marshalPrivKey(privateKey)
}

func RsaDecrypt(key string, secret string) (string, error) {
	ct, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		panic(err.Error())
	}

	privKey, err := unmarshalPrivKey(key)
	if err != nil {
		return "", err
	}

	label := []byte("OAEP Encrypted")
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, privKey, ct, label)
	if err != nil {
		return "", fmt.Errorf("decrypt RSA: %s", err)
	}

	return string(plaintext), nil
}

func marshalPrivKey(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func unmarshalPrivKey(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("parse PEM block")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

// ERRORS
// ====================================================================

// ErrInvalidSecret is returned if a secret of invalid size is used. A secret must be 16, 24 or 32
// bytes exactly.
type ErrInvalidSecret struct {
	secretLen int
}

// Error implements the error interface for ErrInvalidSecret
func (e ErrInvalidSecret) Error() string {
	return fmt.Sprintf("invalid secret: expected length of 16, 24 or 32 bytes, but got %d\n", e.secretLen)
}

func errInvalidSecret(len int) ErrInvalidSecret {
	return ErrInvalidSecret{secretLen: len}
}

// IsInvalidSecretErr is a package helper that returns a boolean indicating wether the provided
// error is an ErrInvalidSecret error.
func IsInvalidSecretErr(err error) bool {
	return errors.As(err, &ErrInvalidSecret{})
}

// ErrInvalidVal is returned if a value fails to Decrypt, this likely means that you never encrypted
// the value with this particular secret in the first place.
var ErrInvalidVal = errors.New("invalid value: your value was likely never encrypted with this service's secret")

func IsInvalidValErr(err error) bool {
	return errors.Is(err, ErrInvalidVal)
}
