package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/Rockup-Consulting/go_std/x/cryptox"
)

const delimiter = "$"

var (
	// ErrInvalidToken is returned whenever a token fails to Decrypt. This can generally mean that
	// a token has been corrupted, but it could also be expired. Generally speaking the difference
	// is irrelevant.
	ErrInvalidToken = errors.New("invalid token")
)

// NewService constructs and returns an auth Service by converting a SecretMap to a pool of
// cryptox.Service(s)
func NewService(secretMap SecretMap) *Service {
	secretKeyToEncryptMap := make(map[uint]cryptox.Service)
	latestSecretKey := uint(0)

	for k, v := range secretMap {
		service, err := cryptox.NewService(v)
		if err != nil {
			panic(err.Error())
		}

		secretKeyToEncryptMap[k] = service

		if k > latestSecretKey {
			latestSecretKey = k
		}
	}

	return &Service{
		mu:                    sync.Mutex{},
		latestSecretKey:       latestSecretKey,
		secretKeyToEncryptMap: secretKeyToEncryptMap,
	}
}

// Service wraps a pool of cryptox.Service(s), making it easier to use encryption in the case of
// multiple secrets, eg. rotating secrets.
type Service struct {
	mu                    sync.Mutex
	latestSecretKey       uint
	secretKeyToEncryptMap map[uint]cryptox.Service
}

// EcryptToken takes some string (token) and encrypts it. The steps taken works as follows:
//   - Encrypt the token
//   - Format the token as: [secret ID]$[encrypted value]
//   - Base64URL Encode the token
//   - Return the encrypted token as string
func (s *Service) EncryptToken(token string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	encryptService, ok := s.secretKeyToEncryptMap[s.latestSecretKey]
	if !ok {
		panic("invalid 'latest secret key', this should not happen, fix it")
	}

	encryptedVal := encryptService.Encrypt(token)

	valueToEncode := fmt.Sprintf("%d%s%s", s.latestSecretKey, delimiter, encryptedVal)
	encodedValue := base64.URLEncoding.EncodeToString([]byte(valueToEncode))

	return encodedValue
}

// UnencryptToken attempts to decrypt a token, if the token cannot be unencrypted, an
// ErrInvalidToken error is returned.
//   - Receive token as string
//   - Base64URL Decode the string
//   - Parse the resulting string as [secret ID]$[encrypted value]
//   - Get encryption service from parse result
//   - Unencrypt the token
//   - Return unencrypted token or error
func (s *Service) UnencryptToken(token string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	unencodedValue, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return "", ErrInvalidToken
	}

	before, after, found := strings.Cut(string(unencodedValue), delimiter)
	if !found {
		return "", ErrInvalidToken
	}

	secretKey, err := strconv.ParseUint(before, 10, 64)
	if err != nil {
		return "", ErrInvalidToken
	}

	encryptService, ok := s.secretKeyToEncryptMap[uint(secretKey)]
	if !ok {
		return "", ErrInvalidToken
	}

	unencryptedToken, err := encryptService.Decrypt(after)
	if err != nil {
		return "", ErrInvalidToken
	}

	return unencryptedToken, nil
}
