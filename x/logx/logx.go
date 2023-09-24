// Package logx provides a set of extensions and convencience functions on top of the built-in log
// package. The standard log.Logger type should remain your main dependency.
package logx

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/Rockup-Consulting/go_std/x/twx"
)

const DefaultFlags = log.Ldate | log.Ltime | log.LUTC | log.Lmsgprefix

// New creates a stdout *log.Logger with our default flags, it takes an optional
// prefix as an argument.
func New(prefix string) *log.Logger {
	if prefix != "" {
		prefix = prefix + ": "
	}

	return log.New(os.Stdout, prefix, DefaultFlags)
}

// NewWriteLogger creates a new *log.Logger that writes to a io.Writer. It takes an
// io.Writer and an optional prefix as arguments.
func NewWriteLogger(w io.Writer, prefix string) *log.Logger {
	if prefix != "" {
		prefix = prefix + ": "
	}

	return log.New(w, prefix, DefaultFlags)
}

// NewDiscard returns a new logger that discards all logs, this is useful when logs are required
// for development, but can be discarded in production
func NewDiscard() *log.Logger {
	return log.New(io.Discard, "", DefaultFlags)
}

// NewTestLogger creates and returns a new *log.Logger with output specifically
// formatted to read nicely in test output. It takes a testing.TB as an argument
// and returns a *log.Logger as well as a function to flush the logs and write
// them to stdout
func NewTestLogger(t testing.TB) (*log.Logger, func()) {
	tw := twx.NewWriter(os.Stdout)

	fmt.Println()
	fmt.Fprintf(os.Stdout, "-------------------- %s LOGS --------------------\n", t.Name())
	l := log.New(tw, "\t", log.Lshortfile|log.Lmsgprefix)

	return l, func() {
		// flush the tabwriter and append an empty line below the logs (for cleaner output)
		tw.Flush()
		fmt.Println()
	}
}
