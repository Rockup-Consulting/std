package cryptox_test

import (
	"encoding/base64"
	"testing"

	"github.com/Rockup-Consulting/go_std/x/cryptox"
)

const (
	TestEncryptSecret     = "thishastobe32bytesforittowork!:)"
	RandomString          = "somerandomstring"
	RandomStringEncrypted = "n9CYLKwNqOY1-M1Kix9de7YhZM17xZ5NTxFh5OL9d1AnVulyRPjr48VSHTg="
)

// Nothing else we can test here
func TestEncrypt(t *testing.T) {
	service, err := cryptox.NewService(TestEncryptSecret)
	isNotErr(t, err)

	encrypted := service.Encrypt(RandomString)

	out := base64.URLEncoding.EncodeToString([]byte(encrypted))

	t.Log(out)
}

func TestDecrypt(t *testing.T) {
	service, err := cryptox.NewService(TestEncryptSecret)
	isNotErr(t, err)

	unencoded, err := base64.URLEncoding.DecodeString(RandomStringEncrypted)
	isNotErr(t, err)

	got, err := service.Decrypt(string(unencoded))
	isNotErr(t, err)

	want := RandomString

	if got != want {
		t.Errorf("RandomString = %s, want = %s", got, want)
	}
}

func TestBothWays(t *testing.T) {
	service, err := cryptox.NewService(TestEncryptSecret)
	isNotErr(t, err)

	encryptedString := service.Encrypt(RandomString)

	got, err := service.Decrypt(encryptedString)
	isNotErr(t, err)

	want := RandomString

	if got != want {
		t.Errorf("got = %s; want = %s", got, want)
	}
}

func isNotErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected nil err, but got %s", err)
	}
}
