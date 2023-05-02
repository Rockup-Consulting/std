// Package u provides some flow control and logic utilities
package u

// Use is a helper that allows us to temporarily silence the compiler about unused variables.
// Sometimes when testing it can be annoying to have to remove all unused variables just to add them
// back in again. Just add them all the this function to silence the compiler.
func Use(vals ...any) {}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

func MustOk[T any](t T, ok bool) T {
	if !ok {
		panic("runtime assertion failed: expected true but got false")
	}

	return t
}
