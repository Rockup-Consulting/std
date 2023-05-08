package cli

import (
	"context"
	"fmt"
)

type Func func(ctx context.Context, args ...string) error

type fn struct {
	f    Func
	name string
	desc string
}

func (f fn) Exec(ctx context.Context, args ...string) error {
	return f.f(ctx, args...)
}

func (f fn) print() string {
	return fmt.Sprintf("\t%s\t%s\n", f.name, f.desc)
}
