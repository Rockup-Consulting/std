// Package form provides helpers for sanitising form input values. The rule is that if an error is
// returned, it should be safe to send directly to the end user in the browser.
package form

import (
	"errors"
	"fmt"

	"strings"
	"time"

	"github.com/Rockup-Consulting/go_std/core/regx"
)

// Field is a form field helper, mostly useful for browser form input
type Field struct {
	Value string
	Err   string
}

type ValidationFunc func(name string, val string) (string, error)

func Validate(name, val string, f ...ValidationFunc) (string, error) {
	sVal := val

	for _, ff := range f {
		x, err := ff(name, sVal)
		if err != nil {
			return sVal, err
		}
		sVal = x
	}

	return sVal, nil
}

var NonEmpty ValidationFunc = func(name, val string) (string, error) {
	if val == "" {
		return val, fmt.Errorf("%q field is required", name)
	}
	return val, nil
}

var Alpha ValidationFunc = func(name, val string) (string, error) {
	ok := regx.MatchAlpha(val)
	if !ok {
		return val, fmt.Errorf("%q field can only have values a-z or A-Z", name)
	}

	return val, nil
}

var Numeric ValidationFunc = func(name, val string) (string, error) {
	ok := regx.MatchNumeric(val)
	if !ok {
		return val, fmt.Errorf("%q field can only have values 0-9", name)
	}

	return val, nil
}

var AlphaNumeric ValidationFunc = func(name, val string) (string, error) {
	ok := regx.MatchAlphaNumeric(val)
	if !ok {
		return val, fmt.Errorf("%q field can only have values 0-9, a-z or A-Z", name)
	}

	return val, nil
}

var Matches = func(m ...string) ValidationFunc {
	return func(name, val string) (string, error) {
		for _, mm := range m {
			if val == mm {
				return val, nil
			}
		}

		return val, fmt.Errorf("%q does not match ", name)
	}
}

var MatchesField = func(fieldName, fieldVal string) ValidationFunc {
	return func(name, val string) (string, error) {
		if val != fieldVal {
			return val, fmt.Errorf("%q does not match field %q ", name, fieldName)
		}
		return val, nil
	}
}

var MinLen = func(min int) ValidationFunc {
	return func(name, val string) (string, error) {
		if len(val) < min {
			return val, fmt.Errorf("%q is too short - min length is %d", name, min)
		}

		return val, nil
	}
}

var MaxLen = func(max int) ValidationFunc {
	return func(name, val string) (string, error) {
		if len(val) > max {
			return val, fmt.Errorf("%q is too long - max length is %d", name, max)
		}

		return val, nil
	}
}

var Len = func(l int) ValidationFunc {
	return func(name, val string) (string, error) {
		if len(val) != l {
			return val, fmt.Errorf("%q should be exactly %d characters long", name, l)
		}

		return val, nil
	}
}

type DateOption struct {
	Val int
	Msg string
}

type DateOptions struct {
	MaxYear  DateOption
	MinYear  DateOption
	MaxMonth DateOption
	MinMonth DateOption
	MaxDay   DateOption
	MinDay   DateOption
}

// Date validates that a date-string is in the format YYYY-MM-DD. If DateOptions are provided, they
// are validated as well.
var Date = func(opts DateOptions) ValidationFunc {
	// no specific reason for these numbers, they just seems useless
	//
	// zero is the no-check value
	if (opts.MinYear.Val < 1900 && opts.MinYear.Val != 0) ||
		opts.MaxYear.Val > 3000 ||
		(opts.MinMonth.Val < 1 && opts.MinMonth.Val != 0) ||
		opts.MaxMonth.Val > 12 ||
		(opts.MinDay.Val < 1 && opts.MinDay.Val != 0) ||
		opts.MaxDay.Val > 31 {

		panic("invalid configuration")
	}

	invalidDate := func(fieldName string) error {
		return fmt.Errorf("%q field is invalid, ensure that date is in format YYYY-MM-DD", fieldName)
	}

	return func(name, val string) (string, error) {
		d, err := time.Parse("2006-01-02", val)
		if err != nil {
			return val, invalidDate(name)
		}

		// YYYY
		if opts.MaxYear.Val != 0 && opts.MaxYear.Val < d.Year() {
			return val, errors.New(opts.MaxYear.Msg)
		}

		if opts.MinYear.Val != 0 && opts.MinYear.Val > d.Year() {
			return val, errors.New(opts.MinYear.Msg)
		}

		// MM
		if opts.MaxMonth.Val != 0 && opts.MaxMonth.Val < int(d.Month()) {
			return val, errors.New(opts.MaxMonth.Msg)
		}

		if opts.MinMonth.Val != 0 && opts.MinMonth.Val > int(d.Month()) {
			return val, errors.New(opts.MinMonth.Msg)
		}

		// DD
		if opts.MaxDay.Val != 0 && opts.MaxDay.Val < d.Day() {
			return val, errors.New(opts.MaxDay.Msg)
		}

		if opts.MinDay.Val != 0 && opts.MinDay.Val > d.Day() {
			return val, errors.New(opts.MinDay.Msg)
		}

		return val, nil
	}
}

var Password = func(reason string, nomatch ...string) ValidationFunc {
	return func(name, val string) (string, error) {
		for _, n := range nomatch {
			if n != "" && strings.Contains(val, n) {
				return val, errors.New(reason)
			}
		}

		if len(val) < 8 {
			return val, fmt.Errorf("%q field is too short, must be atleast 8 characters", name)
		}
		return val, nil
	}
}
