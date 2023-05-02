package cli

import (
	"fmt"
	"io"
)

// Func implements executable
type Func func(args []string) error

type funcExec struct {
	this *command
	f    Func
}

func (f funcExec) Exec(args []string) error {
	return f.f(args)
}

func (m *Menu) createFunc(arg string, d, h string, f Func) *command {
	cmnd := &command{
		a:   m.a,
		arg: arg,
		w:   m.this.w,
		d:   d,
		h:   h,
	}

	e := funcExec{
		this: cmnd,
		f:    f,
	}

	cmnd.e = e

	return cmnd
}

func (f funcExec) Print(w io.Writer) {
	fmt.Fprintf(w, "  %s\t%s\n", f.this.arg, f.this.d)
}
