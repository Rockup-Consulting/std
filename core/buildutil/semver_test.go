package buildutil_test

import (
	"rockup/core/buildutil"
	"testing"
)

func TestSemver(t *testing.T) {
	cases := []struct {
		Start string
		Want  string
		Bump  string
	}{
		{
			Start: "v0.0.0",
			Want:  "v0.0.1",
			Bump:  "patch",
		},
		{
			Start: "v0.0.0",
			Want:  "v0.1.0",
			Bump:  "minor",
		},
		{
			Start: "v0.0.0",
			Want:  "v1.0.0",
			Bump:  "major",
		},
		{
			Start: "v1.9.9",
			Want:  "v1.9.10",
			Bump:  "patch",
		},
		{
			Start: "v1.9.9",
			Want:  "v1.10.0",
			Bump:  "minor",
		},
		{
			Start: "v9.9.9",
			Want:  "v10.0.0",
			Bump:  "major",
		},
	}

	for _, tt := range cases {
		semver := buildutil.BuildSemver(tt.Start)

		if tt.Bump == "patch" {
			semver.PatchBump()
		} else if tt.Bump == "minor" {
			semver.MinorBump()
		} else if tt.Bump == "major" {
			semver.MajorBump()
		} else {
			t.Fatalf("invalid bump type %q", tt.Bump)
		}

		got := semver.String()
		if got != tt.Want {
			t.Errorf("expected %q but got %q", tt.Want, got)
		}
	}
}
