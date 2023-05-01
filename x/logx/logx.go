// Package logx provides a set of extensions and convencience functions on top of the built-in log
// package. The standard log.Logger type should remain your main dependency.
package logx

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/Rockup-Consulting/go_std/x/twx"
)

const defaultFlags = log.Ldate | log.Ltime | log.LUTC | log.Lmsgprefix

// New creates a stdout *log.Logger with our default flags, it takes an optional
// prefix as an argument.
func New(prefix string) *log.Logger {
	if prefix != "" {
		prefix = prefix + ": "
	}

	return log.New(os.Stdout, prefix, defaultFlags)
}

// NewWriteLogger creates a new *log.Logger that writes to a io.Writer. It takes an
// io.Writer and an optional prefix as arguments.
func NewWriteLogger(w io.Writer, prefix string) *log.Logger {
	if prefix != "" {
		prefix = prefix + ": "
	}

	return log.New(w, prefix, defaultFlags)
}

// NewBufferedLogger creates a new *log.Logger that writes to a buffer. It
// takes an optional prefix as an argument and returns a *log.Logger as well
// as a function to flush the logger.
func NewBufferedLogger(prefix string) (*log.Logger, func() string) {
	if prefix != "" {
		prefix = prefix + ": "
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	l := log.New(w, prefix, defaultFlags)

	return l, func() string {
		w.Flush()
		return buf.String()
	}
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
