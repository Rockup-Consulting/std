// package v exposes useful validation functions and parsing functions
package v

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func IsEmpty(val string) bool {
	if strings.TrimSpace(val) == "" {
		return true
	}

	return false
}

func IsNotEmpty(val string) bool {
	if strings.TrimSpace(val) == "" {
		return false
	}

	return true
}

func IsValidEmail(val string) bool {
	if len(val) > 254 || !rxEmail.MatchString(val) {
		return false
	}

	return true
}

func IsMoreCharactersThan(val string, min int) bool {
	l := utf8.RuneCountInString(val)
	if l < min {
		return false
	}

	return true
}

func IsLessCharactersThan(val string, max int) bool {
	l := utf8.RuneCountInString(val)
	if l > max {
		return false
	}

	return true
}

func IsURL(val string) (*url.URL, bool) {
	u, err := url.Parse(val)
	if err != nil {
		return nil, false
	} else if u.Scheme == "" || u.Host == "" {
		return nil, false
	} else if u.Scheme != "http" && u.Scheme != "https" {
		return nil, false
	}

	return u, true
}

func IsInteger(val string) (int, bool) {
	n, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}

	return n, true
}

func IsFloat(val string) (float64, bool) {
	n, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, false
	}

	return n, true
}

// ====================================================================
// Utils

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
