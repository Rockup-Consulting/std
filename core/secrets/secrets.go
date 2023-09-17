package secrets

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// SecretMap is simply a mapping from uint -> string. Secret ID's should be incremented by one.
type Map map[uint64]string

// SecretMapFromJSON accepts a JSON string as a parameter and returns a key->val map, of type
// uint->string, and an error.
//
// Requirements:
//   - secret should be exactly 16/24/32bytes
//   - key should be a uint (preferrably incremented by one from the previous key). Easiest way to
//     generate a secret is to use rand.
func MapFromJSON(jsonStr string) (Map, error) {
	secretsMap := make(Map)
	err := json.Unmarshal([]byte(jsonStr), &secretsMap)

	if err != nil {
		return nil, err
	}

	for _, v := range secretsMap {
		valueByteLength := len([]byte(v))
		if valueByteLength != 32 {
			return nil, fmt.Errorf("parse secret: expected the secret to be 32 bytes, but got: %dbytes", valueByteLength)
		}
	}

	return secretsMap, nil
}

// SecretMapFromBase64String accepts a string as a parameter and returns a
// key->val map, of type uint->string, and an error.
//
// Requirements:
//   - argument should be a base64URL encoded string of a JSON document that was a uint->32byte
//     string map
//   - secrets should be exactly 32bytes
//   - key should be a uint (preferrably incremented by one from the previous key)
func MapFromBase64String(encodedSecrets string) (Map, error) {
	result, err := base64.URLEncoding.DecodeString(encodedSecrets)
	if err != nil {
		return nil, err
	}

	return MapFromJSON(string(result))
}
