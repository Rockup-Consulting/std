package buildutil

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/mod/semver"
)

type semverutil struct{ string }

func BuildSemver(version string) semverutil {

	if !semver.IsValid(version) {
		panic("invalid semver")
	}

	return semverutil{version}
}

func (s *semverutil) MajorBump() {
	s.prep(func(v [3]int) [3]int {
		v[0]++
		v[1] = 0
		v[2] = 0
		return v
	})
}

func (s *semverutil) MinorBump() {
	s.prep(func(v [3]int) [3]int {
		v[1]++
		v[2] = 0
		return v
	})
}

func (s *semverutil) PatchBump() {
	s.prep(func(v [3]int) [3]int {
		v[2]++
		return v
	})
}

func (s *semverutil) prep(f func(v [3]int) [3]int) {
	temp := strings.TrimPrefix(s.string, "v")
	slice := strings.Split(temp, ".")
	if len(slice) != 3 {
		panic("invalid semver, expected major, minor and patch values")
	}

	var intSlice [3]int
	for i, x := range slice {
		val, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}

		intSlice[i] = val
	}

	out := f(intSlice)

	outStr := "v"
	for i, x := range out {
		if i == 0 {
			outStr = fmt.Sprintf("%s%d", outStr, x)
		} else {
			outStr = fmt.Sprintf("%s.%d", outStr, x)
		}
	}

	s.string = outStr
}

func (s semverutil) String() string {
	return string(s.string)
}
