package auth_test

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/Rockup-Consulting/go_std/core/auth"
)

const (
	SessionID     = "bfe7a9cb-827b-493b-9c2a-111111111111"
	RememberToken = "rememberme"
	userID        = "bfe7a9cb-827b-493b-9c2a-4f449b98f89a"
)

var (
	SecretMap = map[uint]string{
		1: "thishastobe32bytesforittowork!:)",
		2: "thishastobe32aswellforittowork:)",
	}

	wantMap = auth.SecretMap{
		1: "thishastobe32bytesforittowork!:)",
		2: "thishastobe32aswellforittowork:)",
	}

	AuthURL = url.URL{
		Scheme: "http",
		Host:   "localhost:3000",
	}
)

func TestSecretMapFromJson(t *testing.T) {
	jsonStr := `{
		"1": "thishastobe32bytesforittowork!:)",
		"2": "thishastobe32aswellforittowork:)"
}`

	secretMap, err := auth.SecretMapFromJSON(jsonStr)
	if err != nil {
		t.Fatalf("expected nil error but got %s", err)
	}

	if !reflect.DeepEqual(secretMap, wantMap) {
		t.Errorf("auth.SecretMapFromBase64String(str) = %v; want %v", secretMap, wantMap)
	}
}

func TestSecretMapFromBase64String(t *testing.T) {
	confStr := "ewogICAgIjEiOiAidGhpc2hhc3RvYmUzMmJ5dGVzZm9yaXR0b3dvcmshOikiLAogICAgIjIiOiAidGhpc2hhc3RvYmUzMmFzd2VsbGZvcml0dG93b3JrOikiCn0="
	secretMap, err := auth.SecretMapFromBase64String(confStr)

	if err != nil {
		t.Fatalf("expected nil error but got %s", err)
	}

	if !reflect.DeepEqual(secretMap, wantMap) {
		t.Errorf("auth.SecretMapFromBase64String(str) = %v; want %v", secretMap, wantMap)
	}
}
