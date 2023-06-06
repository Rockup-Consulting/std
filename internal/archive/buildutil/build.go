package buildutil

import (
	"fmt"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

// BuildGoBin builds a Go binary in the path "{projectRoot}/bin/{orgName}/{os_arch}/{appName}"
// The build target is for linux since our Docker images run on Alpine.
func BuildGoBin(info Info) error {
	binPath := "./" + filepath.Join("bin", info.OrgName, info.Os+"_"+info.Arch, info.AppName)
	srcPath := "./" + filepath.Join("orgs", info.OrgName, "services", info.AppName)

	// remove binary (without printing result to shell)
	_, _ = sh.Exec(nil, nil, nil, "rm", binPath)

	fmt.Printf("building %s %s %s %s binary\n", info.OrgName, info.AppName, info.Os, info.Arch)

	err := sh.Run(
		"env", "GOOS="+info.Os, "GOARCH="+info.Arch, "CGO_ENABLED=0",
		"go",
		"build", "-o", binPath,
		srcPath,
	)

	if err != nil {
		return err
	}

	return nil
}
