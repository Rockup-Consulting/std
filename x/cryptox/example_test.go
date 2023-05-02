package cryptox_test

import (
	"fmt"

	"github.com/Rockup-Consulting/go_std/x/cryptox"
)

func ExampleService() {

	// Try to create cryptox.Service with invalid secret, you will receive an error of type
	// ErrInvalidSecret
	invalidSecret := "invalid_secret_not_16_24_or_32_bytes"

	_, err := cryptox.NewService(invalidSecret)
	if err != nil && cryptox.IsInvalidSecretErr(err) {
		fmt.Println("invalid secret is invalid")
	}

	// Create cryptox.Service using a valid secret, this will succeed

	secret := "super_secret_value_of_32byte_len"
	encService, err := cryptox.NewService(secret)

	if err != nil {
		panic(err.Error())
	}

	// Encrypt and Decrypt some value

	myVal := "supercalifragilisticexpialidocious"

	encryptedVal := encService.Encrypt(myVal)
	unencryptedVal, err := encService.Decrypt(encryptedVal)
	if err != nil {
		panic(err.Error())
	}

	// Compare encrypted value to unencrypted value

	fmt.Printf("myVal == unencryptedVal => %t\n", myVal == unencryptedVal)

	// try to unencrypt a value that was not encrypted by our secret

	invalidVal := "this has not been encrypted by my secret"
	_, err = encService.Decrypt(invalidVal)
	if err != nil && cryptox.IsInvalidValErr(err) {
		fmt.Println("invalid encrypted value is invalid")
	}

	// Output:
	// invalid secret is invalid
	// myVal == unencryptedVal => true
	// invalid encrypted value is invalid

}
