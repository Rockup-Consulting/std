package buildutil

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func RunTests() bool {
	fmt.Println("running tests")

	err := sh.Run("go", "test", "./...")
	if err != nil {
		fmt.Printf("tests did not pass: %s\n", err)
		return false
	}

	return true
}
