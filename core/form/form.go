// package form is a utility package for working with HTML forms
package form

type State int

const (
	Initial State = iota
	Success
	Error
)

// Field represents a form field
type Field struct {
	Val string
	Err string
}

// Val returns a form.Field with just a value set
func Val(val string) Field {
	return Field{
		Val: val,
	}
}
