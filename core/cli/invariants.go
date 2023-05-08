package cli

import (
	"fmt"
	"strings"
)

func cliInitErr(s string) error {
	return fmt.Errorf("CLI Init Error: %s", s)
}

func validateName(name string) {
	if strings.Contains(name, " ") {
		panic(cliInitErr(fmt.Sprintf("invalid arg name %q, spaces not allowed", name)))
	}
}
