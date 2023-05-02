package auth_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"rockup/core/auth"

	"github.com/Rockup-Consulting/go_std/x/randx"
)

func ExampleService() {

	// This first part is just to demonstrate the setup, the base64Secrets string should obviously
	// come from outside your application, eg. ENV variable
	jsonSecrets := `{
		"1": "thishastobe32bytesforittowork!:)",
		"2": "thishastobe32aswellforittowork:)"
}`
	base64Secrets := base64.URLEncoding.EncodeToString([]byte(jsonSecrets))

	// setup auth service

	secretMap, err := auth.SecretMapFromBase64String(base64Secrets)
	if err != nil {
		panic(err.Error())
	}

	authService := auth.NewService(secretMap)

	// dummy auth session token
	type Token struct {
		Name          string
		SecurityToken string
	}

	token, err := randx.UID()
	if err != nil {
		panic(err.Error())
	}

	authSessionToken := Token{
		Name:          "ferdinand",
		SecurityToken: token,
	}

	// use auth service

	sessionStr, err := json.Marshal(authSessionToken)
	if err != nil {
		panic(err.Error())
	}

	encryptedToken := authService.EncryptToken(string(sessionStr))

	unencryptedToken, err := authService.UnencryptToken(encryptedToken)
	if err != nil {
		panic(err.Error())
	}

	// unmarshal string back into session
	unmarshaledAuthSessionToken := Token{}

	err = json.Unmarshal([]byte(unencryptedToken), &unmarshaledAuthSessionToken)
	if err != nil {
		panic(err.Error())
	}

	// check that everything worked
	fmt.Println(string(sessionStr) == unencryptedToken)
	fmt.Println(reflect.DeepEqual(authSessionToken, unmarshaledAuthSessionToken))

	// Output:
	// true
	// true
}
