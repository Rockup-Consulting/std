package cryptox

import (
	"errors"
	"fmt"
)

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
