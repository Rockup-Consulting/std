package dotenv

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func Set() error {
	env, err := os.Open(".env")
	if err != nil {
		return err
	}

	bufreader := bufio.NewReader(env)

	var eof bool
	for !eof {
		l, err := bufreader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				eof = true
			} else {
				return err
			}
		}

		// check to see if line is commented out
		if strings.HasPrefix(l, "# ") {
			continue
		}

		if strings.Contains(l, " ") {
			return fmt.Errorf("invalid env line, spaces are not allowed: %q", l)
		}

		name, val, ok := strings.Cut(l, "=")
		if !ok {
			return fmt.Errorf("invalid env line: %q", l)
		}

		if val[len(val)-1] == '\n' {
			val = val[:len(val)-1]
		}

		fmt.Printf("setting %q\n", name)
		os.Setenv(name, val)
	}

	return nil
}
