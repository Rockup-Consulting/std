package validate

import (
	"errors"
	"fmt"
)

type FailureReason string

const (
	TooShort             FailureReason = "TOO_SHORT"
	TooLong              FailureReason = "TOO_LONG"
	TooSmall             FailureReason = "TOO_SMALL"
	TooLarge             FailureReason = "TOO_LARGE"
	Invalid              FailureReason = "INVALID"
	InvalidFormat        FailureReason = "INVALID_FORMAT"
	MissingRequiredField FailureReason = "MISSING_REQUIRED_FIELD"
	IncorrectType        FailureReason = "INCORRECT_TYPE"
	FailedToParse        FailureReason = "FAILED_TO_PARSE"
	ShouldBeNumeric      FailureReason = "SHOULD_BE_NUMERIC"
	Mismatch             FailureReason = "MISMATCH"
	DoesNotExist         FailureReason = "DOES_NOT_EXIST"
)

func EitherOrFailure(fields ...string) FailureReason {
	msg := "REQUIRES FIELDS"

	for i, f := range fields {
		if i == 0 {
			msg = fmt.Sprintf("%s %s", msg, f)
			continue
		}
		msg = fmt.Sprintf("%s OR %s", msg, f)
	}

	return FailureReason(msg)
}

// NewError creates an instance or validation.Error with initialised values.
func NewError() Error {
	return Error{
		ordered: make([]FieldError, 0),
		indexed: make(map[string]FieldError),
	}
}

type FieldError struct {
	Field  string
	Reason FailureReason
}

type Error struct {
	// the slice maintains the error order for cleaner log output
	ordered []FieldError
	// the map allows easy access (eg. in templates)
	indexed map[string]FieldError
}

// Error implements the error interface for validation.Error
func (e Error) Error() string {
	var fields string
	for i, f := range e.ordered {
		if i == 0 {
			fields = fmt.Sprintf("%q", f.Field)
			continue
		}
		fields += fmt.Sprintf(", %q", f.Field)
	}

	return fmt.Sprintf("validation errors on fields: %s", fields)
}

// Get returns the value corresponding to the given key along with a boolean that is true if
// a value was found and false if there were no key with the given name.
func (e Error) Get(field string) (FieldError, bool) {
	fe, ok := e.indexed[field]
	return fe, ok
}

// Count returns the number of FieldErrors available on the ErrValidation instance.
func (e Error) Count() int {
	return len(e.ordered)
}

func (e *Error) CreateSetField(field string) func(reason FailureReason) {
	return func(r FailureReason) {
		e.SetField(field, r)
	}
}

// SetField adds a FieldError to the validation.Error. If the field has already been set, the
// previous one is silently overwritten. Note that the new FieldError will be appended to the end
// of the list.
func (e *Error) SetField(field string, reason FailureReason) {

	if _, alreadyExists := e.indexed[field]; alreadyExists {
		// remove from the slice

		for i, ee := range e.ordered {
			if ee.Field == field {
				e.ordered = append(e.ordered[:i], e.ordered[i+1:]...)
				break
			}
		}
	}

	fe := FieldError{field, reason}
	e.indexed[field] = fe
	e.ordered = append(e.ordered, fe)
}

func (e Error) Errors() []FieldError {
	return e.ordered
}

// IsErr reports whether any fields on validation.Error has been set. This, of course, implies that
// that the error is in fact an error.
func (e Error) IsErr() bool {
	return len(e.ordered) > 0
}

// IsError reports whether any error in err's chain matches validation.Error
func IsError(err error) bool {
	return errors.Is(err, &Error{})
}

// GetError finds the first error in err's chain that matches validation.Error, and if one is found,
// returns that error value and true. Otherwise, it returns false.
func GetError(err error) (Error, bool) {
	var ee Error
	if !errors.As(err, &ee) {
		return Error{}, false
	}

	return ee, true
}
