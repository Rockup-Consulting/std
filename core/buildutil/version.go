package buildutil

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/magefile/mage/sh"
)

type Info struct {
	AppName       string
	OrgName       string
	Os            string //default is linux
	Arch          string //default is amd64
	Version       string
	BuildHash     string
	BuildTimeUTC  string
	BuildTimeSAST string
}

var (
	//go:embed appName
	appName string

	//go:embed orgName
	orgName string

	//go:embed buildHash
	buildHash string

	//go:embed buildVersion
	buildVersion string

	//go:embed buildTimeUTC
	buildTimeUTC string

	//go:embed buildTimeSAST
	buildTimeSAST string

	//go:embed buildOs
	buildOs string

	//go:embed buildArch
	buildArch string

	InfoEmbed = Info{
		OrgName:       orgName,
		AppName:       appName,
		Version:       buildVersion,
		BuildHash:     buildHash,
		BuildTimeUTC:  buildTimeUTC,
		BuildTimeSAST: buildTimeSAST,
		Os:            buildOs,
		Arch:          buildArch,
	}
)

func buildTimeCheck(packageDir string) {
	dir, err := os.Stat(packageDir)
	if err != nil {
		panic(err)
	}

	if !dir.IsDir() {
		panic("not a directory")
	}
}

func NewInfo(orgName, appName, version, os, arch string) (Info, error) {
	buildHash, err := sh.Output("git", "rev-parse", "--short", "HEAD")
	if err != nil {
		return Info{}, err
	}

	now := time.Now()
	format := "2006-01-02 15:04:05 MST"
	buildTimeUTC := now.UTC().Format(format)
	buildTimeSAST := now.Local().Format(format)

	osStr := "linux"
	archStr := "amd64"

	if os != "" {
		osStr = os
	}

	if arch != "" {
		archStr = arch
	}

	info := Info{
		AppName:       appName,
		OrgName:       orgName,
		Os:            osStr,
		Arch:          archStr,
		Version:       version,
		BuildHash:     buildHash,
		BuildTimeUTC:  buildTimeUTC,
		BuildTimeSAST: buildTimeSAST,
	}

	return info, nil
}

// BuildFunc refers to a function that is passed to EmbedBuildInfo. The application binary must be
// built in the BuildFunc for the Info to be embedded in the application
type BuildFunc func(info Info) error

// Generate is a BUILDTIME function that generates the necessary build info and then executes the
// BuildFunc(s). After the BuildFunc(s) have been executed, the build info is replaced with dummy
// information. The binary/image build must happen within the BuildFunc(s) for the build information
// to be embedded in the executable.
func EmbedBuildInfo(i Info, f ...BuildFunc) error {
	_, callerPath, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("problem getting package path")
	}

	packageDir, _ := filepath.Split(callerPath)

	buildTimeCheck(packageDir)

	if err := writeFile("appName", i.AppName, packageDir); err != nil {
		return err
	}
	if err := writeFile("buildOs", i.Os, packageDir); err != nil {
		return err
	}
	if err := writeFile("buildArch", i.Arch, packageDir); err != nil {
		return err
	}
	if err := writeFile("orgName", i.OrgName, packageDir); err != nil {
		return err
	}
	if err := writeFile("buildHash", i.BuildHash, packageDir); err != nil {
		return err
	}
	if err := writeFile("buildVersion", i.Version, packageDir); err != nil {
		return err
	}
	if err := writeFile("buildTimeUTC", i.BuildTimeUTC, packageDir); err != nil {
		return err
	}
	if err := writeFile("buildTimeSAST", i.BuildTimeSAST, packageDir); err != nil {
		return err
	}

	// execute buildfunc
	for _, ff := range f {
		err := ff(i)
		if err != nil {
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	// replace build info with dummy info
	if err := writeFile("appName", "no name", packageDir); err != nil {
		return err
	}
	if err := writeFile("orgName", "mega corp", packageDir); err != nil {
		return err
	}
	if err := writeFile("buildOs", "", packageDir); err != nil {
		return err
	}
	if err := writeFile("buildArch", "", packageDir); err != nil {
		return err
	}
	if err := writeFile("buildHash", "hash", packageDir); err != nil {
		return err
	}
	if err := writeFile("buildVersion", "develop", packageDir); err != nil {
		return err
	}
	if err := writeFile("buildTimeUTC", "now", packageDir); err != nil {
		return err
	}
	if err := writeFile("buildTimeSAST", "now", packageDir); err != nil {
		return err
	}

	return nil
}

// utility for writing information to the specified file
func writeFile(fileName, content, packageDir string) error {
	fp := filepath.Join(packageDir, fileName)

	f, err := os.Create(fp)
	if err != nil {
		return err
	}

	fmt.Fprint(f, content)

	return nil
}

// Version Bumping
//
// All applications are versioned using either Semver or BuildNumber. We track application versions
// using git tags in the format of:
// orgName/appName-[semver/buildnum] | rockup/http_api-1.0.9

const (
	Patch string = "patch"
	Minor string = "minor"
	Major string = "major"
)

func (g *GitClient) BumpSemver(orgName, appName string, bump string) (string, bool) {
	fmt.Println("Bumping version")

	// get latest git tag where tag format is: [orgName]/[appName]-[semver]
	iter, err := g.repo.Tags()
	if err != nil {
		panic(err)
	}

	version := "v0.0.0"
	head, err := g.repo.Head()
	if err != nil {
		panic(err)
	}

	hash := head.Hash()

	for {
		t, err := iter.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err.Error())
		}

		fmt.Println(t)
		tag := t.Name().Short()
		org, rest, ok := strings.Cut(tag, "/")
		if !ok {
			panic("invalid version tag")
		}

		if orgName != org {
			continue
		}

		app, semver, ok := strings.Cut(rest, "-")
		if !ok {
			// if the version tag is invalid, we just keep going. This is a way of allowing tags
			// other than versions, like if we want to leave certain comments
			continue
		}

		if app != appName {
			continue
		}

		version = semver
	}

	// parse semver
	smvr := BuildSemver(version)

	// bump
	switch bump {
	case Major:
		smvr.MajorBump()
	case Minor:
		smvr.MinorBump()
	case Patch:
		smvr.PatchBump()
	default:
		panic(fmt.Sprintf("invalid semver bump: %s", bump))
	}

	// tag commit
	versionTag := s(orgName, "/", appName, "-", smvr.String())
	_, err = g.repo.CreateTag(versionTag, hash, nil)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	// push
	err = g.repo.Push(nil)
	if err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			panic("expected a git tag to be pushed")
		} else {
			panic(err.Error())
		}
	}

	return smvr.String(), true
}

func (g *GitClient) BumpBuildNumber() error {
	return nil
}

func AssertSemverBump(bump string) {
	switch bump {
	case Patch:
		return
	case Minor:
		return
	case Major:
		return
	default:
		panic(fmt.Sprintf("Invalid version bump %q. Expected 'major', 'minor' or 'patch'", bump))
	}
}
