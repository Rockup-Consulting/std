package cli

import "fmt"

type ErrAttemptToOverwriteArg struct {
	arg      string
	menuName string
}

func (e ErrAttemptToOverwriteArg) Error() string {
	return fmt.Sprintf(
		`err: attempted to overwrite arg '%s' on menu '%s'

This is a massive bug-surface-area, and so we simply made it illegal. A user cannot overwrite their
own arg on a cliApp. Once you set it, it is set forever.
	`, e.arg, e.menuName)
}

var errAttemptToOverwriteArg = func(arg, menuName string) ErrAttemptToOverwriteArg {
	return ErrAttemptToOverwriteArg{arg, menuName}
}

type ErrAttemptToAddMultiMenuToApp struct{}

func (e ErrAttemptToAddMultiMenuToApp) Error() string {
	return `err: attempted to add multiple menus to app

The App Menu can only be initialized once.
`
}

var errAttemptToReInitMenu = func() ErrAttemptToAddMultiMenuToApp {
	return ErrAttemptToAddMultiMenuToApp{}
}

// ===========

type ErrAttemptToCreateExecWithNilFunc struct {
	key string
}

func (e ErrAttemptToCreateExecWithNilFunc) Error() string {
	return fmt.Sprintf(
		`err: attempted to create an Executable Command with a nil arg - '%s'

This is a useless case and is not allowed.
`, e.key)
}

var errAttemptToCreateExecWithNilFunc = func(key string) ErrAttemptToCreateExecWithNilFunc {
	return ErrAttemptToCreateExecWithNilFunc{key}
}

// ===========

type ErrCmndNotFound struct {
	key string
}

func (e ErrCmndNotFound) Error() string {
	return fmt.Sprintf("err: command not found '%s'", e.key)
}

var errCmndNotFound = func(key string) ErrCmndNotFound {
	return ErrCmndNotFound{key}
}

// ===========

type ErrInternalNilExec struct {
	keys []string
}

func (e ErrInternalNilExec) Error() string {
	return fmt.Sprintf(`err: internal error - reached impossible case

Command created without executable, this should not be possible.
	
Args - %s
`, e.keys)
}

var errInternalNilExec = func(keys []string) ErrInternalNilExec {
	return ErrInternalNilExec{keys}
}

// ===========

type ErrInternal struct{}

func (e ErrInternal) Error() string {
	return "err: internal error - reached impossible state"
}

var errInternal = func() ErrInternal {
	return ErrInternal{}
}

// ===========

type ErrInternalNilGroupsSlice struct{}

func (e ErrInternalNilGroupsSlice) Error() string {
	return `err: internal error - nil Groups slice on Menu

The Groups slice on Menu cannot ever be nil. This is because as soon as a Menu has been created, we
append the standard commands, meaning that there will always be at least one group.
`
}

var errInternalNilGroupsSlice = func() ErrInternalNilGroupsSlice {
	return ErrInternalNilGroupsSlice{}
}

// ===========

type ErrEmptyArg struct {
	menuName string
}

func (e ErrEmptyArg) Error() string {
	return fmt.Sprintf("err: cannot create command with empty key on menu '%s'\n", e.menuName)
}

var errEmptyArg = func(menuName string) ErrEmptyArg {
	return ErrEmptyArg{menuName: menuName}
}

// ===========

type ErrAppMenuNotInitialized struct{}

func (e ErrAppMenuNotInitialized) Error() string {
	return "err: app menu not initialized"
}

var errAppMenuNotInitialized = func() ErrAppMenuNotInitialized {
	return ErrAppMenuNotInitialized{}
}

// ===========

type ErrAttemptToUseReservedKey struct {
	key string
}

func (e ErrAttemptToUseReservedKey) Error() string {
	return fmt.Sprintf("err: attempted to use a reserved key: '%s'\n", e.key)
}

var errAttemptToUseReservedKey = func(key string) ErrAttemptToUseReservedKey {
	return ErrAttemptToUseReservedKey{key}
}

// ===========

type ErrTooManyArgsProvided struct {
	arg string
}

func (e ErrTooManyArgsProvided) Error() string {
	return fmt.Sprintf("err: too many args provided for command '%s'\n", e.arg)
}

var errTooManyArgsProvided = func(arg string) ErrTooManyArgsProvided {
	return ErrTooManyArgsProvided{arg}
}
