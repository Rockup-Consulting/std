package web

import "errors"

// Error is meant to be used as an early return signal. If you want to build functionality with an
// early return, you can simply return a web.Error up the call chain. At the point of needing to
// handle an error, you can choose to do nothing. It is up to the caller to implement the "ignore"
// functionality. This would typically happen in an outer middleware.
var Error = errors.New("web error")

// IsError checks whether an error is of type web.Error
func IsError(err error) bool {
	return errors.Is(err, Error)
}
