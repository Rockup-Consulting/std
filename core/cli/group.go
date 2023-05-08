package cli

import (
	"context"
	"fmt"
)

type group string

// Exec is a dummy method implemented to comply with the Executable interface
func (g group) Exec(ctx context.Context, args ...string) error {
	return nil
}

func (g group) print() string {
	return fmt.Sprintf("\n%s:\n", g)
}
