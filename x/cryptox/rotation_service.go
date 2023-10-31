package cryptox

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/Rockup-Consulting/std/core/secrets"
)

type RotationService struct {
	EncServices map[uint64]Service
	Latest      uint64
}

func NewRotationService(secrets secrets.Map) (RotationService, error) {
	if len(secrets) < 1 {
		return RotationService{}, errors.New("cannot create a RotationService with no secrets")
	}

	out := make(map[uint64]Service, len(secrets))

	latest := uint64(0)
	for id, secret := range secrets {
		enc, err := NewService(secret)
		if err != nil {
			return RotationService{}, err
		}

		if id > latest {
			latest = id
		}

		out[id] = enc
	}

	return RotationService{
		EncServices: out,
		Latest:      latest,
	}, nil
}

var delim = []byte("$")

func (r RotationService) Encrypt(val []byte) []byte {
	enc, ok := r.EncServices[r.Latest]
	if !ok {
		panic(fmt.Sprintf("failed to find latest enc service with ID -> %d", r.Latest))
	}

	encrypted := enc.Encrypt(val)

	out := bytes.Join([][]byte{
		[]byte(fmt.Sprintf("%d", r.Latest)),
		encrypted,
	}, delim)

	return out
}

func (r RotationService) Decrypt(val []byte) ([]byte, bool, error) {
	out := bytes.SplitN(val, delim, 2)
	if len(out) != 2 {
		return nil, false, ErrInvalidVal
	}

	idBytes := out[0]
	encryptedVal := out[1]

	id, err := strconv.ParseUint(string(idBytes), 10, 0)
	if err != nil {
		return nil, false, ErrInvalidVal
	}

	enc, ok := r.EncServices[id]
	if !ok {
		return nil, false, ErrInvalidVal
	}

	decryptedVal, _, err := enc.Decrypt(encryptedVal)
	if err != nil {
		return nil, false, ErrInvalidVal
	}

	var shouldRefresh bool
	if id != r.Latest {
		shouldRefresh = true
	}

	return decryptedVal, shouldRefresh, nil
}
