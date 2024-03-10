package v

import "errors"

var Error = errors.New("validation error")

func IsError(err error) bool {
	return errors.Is(err, Error)
}
