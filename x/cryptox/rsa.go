package cryptox

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

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
