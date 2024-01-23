package randx

import (
	"crypto/rand"
	"math/big"
)

const PINLen = 6
const TokenLen = 32

func String(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(err)
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}

func UID() string {
	return String(32)
}

// Pin generates and returns a random 6 character Pin.
func PIN() string {
	return String(PINLen)
}
