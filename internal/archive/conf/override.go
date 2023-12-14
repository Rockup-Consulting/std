package conf

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

var (
	errOverrideNotProvided = errors.New("no override value provided")

	errOverrideParseFailed = func(key string, field Field, newVal any) error {
		gotVal := newVal
		if field.Sensitive {
			gotVal = "**SENSITIVE_FIELD**"
		}

		return fmt.Errorf("parse config value for %q, expected type %q, got value %q", key, field.Type, gotVal)
	}
)

type Override interface {
	Name() string

	// converts a Field Key to a specific override. eg. if a field has the key
	// ["hello", "world"] then the envOverride would be PREFIX_HELLO_WORLD and the
	// cliOverride would be --prefix-hello-world
	Key(prefix string, field Field) string

	// Optional sanity check method
	Cleanup() error

	// GetSTRING attempts to get a string from the Override, if it fails, an error is returned
	GetSTRING(key string, field Field, overrideKey string) (string, error)

	// GetINT attempts to get an int from the Override, if it fails, an error is returned
	GetINT(key string, field Field, overrideKey string) (int, error)

	// GetDURATION attempts to get a time.Duration from the Override, if it fails, an error is
	// returned.
	//
	// The parsing function is time.ParseDuration, meaning that most formats would work:
	// 1m20s, 10s, 500ms etc. Check [time.ParseDuration]: https://pkg.go.dev/time#ParseDuration for
	// more information.
	GetDURATION(key string, field Field, overrideKey string) (time.Duration, error)

	// GetURL attempts to get a url.URL from the Override, if it fails, an error is
	// returned.
	//
	// The parsing function is url.Parse
	GetURL(key string, field Field, overrideKey string) (url.URL, error)

	// GetBOOL attempts to get a boolean from the Override, if it fails, an error is returned
	GetBOOL(key string, field Field, overrideKey string) (bool, error)
}

// override implementations
var _ Override = (*envOverride)(nil)
