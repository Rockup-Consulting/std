package conf

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type envOverride struct{}

func ENV() *envOverride {
	return &envOverride{}
}

func (c envOverride) Name() string {
	return "ENV_OVERRIDE"
}

func (c envOverride) Key(prefix string, field Field) string {
	keysCount := len(field.Keys.Slice())

	var sb strings.Builder
	for i, k := range field.Keys.Slice() {
		fmt.Fprint(&sb, strings.ToUpper(k))

		if i == keysCount-1 {
			break
		}
		fmt.Fprint(&sb, "_")
	}

	var key string
	if prefix == "" {
		key = sb.String()
	} else {
		key = strings.ToUpper(prefix) + "_" + sb.String()
	}

	return key
}

func (c envOverride) GetSTRING(key string, field Field, overrideKey string) (string, error) {
	env := os.Getenv(overrideKey)
	if env == "" {
		return "", errOverrideNotProvided
	}

	return env, nil
}

func (c envOverride) GetINT(key string, field Field, overrideKey string) (int, error) {
	env := os.Getenv(overrideKey)
	if env == "" {
		return 0, errOverrideNotProvided
	}

	out, err := strconv.Atoi(env)
	if err != nil {
		return 0, errOverrideParseFailed(key, field, env)
	}

	return out, nil
}

func (c envOverride) GetDURATION(key string, field Field, overrideKey string) (time.Duration, error) {
	env := os.Getenv(overrideKey)
	if env == "" {
		return 0, errOverrideNotProvided
	}

	out, err := time.ParseDuration(env)
	if err != nil {
		return 0, errOverrideParseFailed(key, field, env)
	}

	return out, nil
}

func (c envOverride) GetURL(key string, field Field, overrideKey string) (url.URL, error) {
	env := os.Getenv(overrideKey)
	if env == "" {
		return url.URL{}, errOverrideNotProvided
	}

	out, err := url.Parse(env)

	if err != nil {
		return url.URL{}, errOverrideParseFailed(key, field, env)
	}

	return *out, nil
}

func (c envOverride) GetBOOL(key string, field Field, overrideKey string) (bool, error) {
	env := os.Getenv(overrideKey)
	if env == "" {
		return false, errOverrideNotProvided
	}

	switch env {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, errOverrideParseFailed(key, field, env)
	}

}

func (c envOverride) Cleanup() error {
	return nil
}
