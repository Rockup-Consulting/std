package region_test

import (
	"testing"

	"github.com/Rockup-Consulting/go_std/core/region"
)

func TestSanitiseZA(t *testing.T) {

	t.Run("valid numbers", func(t *testing.T) {
		want := "+27836532269"

		testCases := []string{
			"+27836532269",
			"27836532269",
			"0836532269",
			"270836532269",
			"+270836532269",
			"083 653 2269",
			"+2783 653 2269",
			"+27 083 653 2269",
			"+27 83 653 2269",
			"+27 083 6 53 2 269",
		}

		za, ok := region.IsoRegionMap["ZA"]
		if !ok {
			t.Errorf("mobile.IsoCountyMap[ZA] = _, %t; want = %t", ok, true)
		}

		for _, tt := range testCases {
			t.Run(tt, func(t *testing.T) {
				out, err := za.SanitiseMobileNumber(tt)
				if err != nil {
					t.Errorf("mobile.Sanitise() = _, %s; want = nil", err)
				}

				if out != want {
					t.Errorf("mobile.Sanitise() = %v; want = %v", out, want)
				}
			})
		}
	})

	t.Run("invalid numbers", func(t *testing.T) {
		testCases := []string{
			"+278365322699",
			"+2783653226",
			"08365322699",
			"083653226",
			"+27083653226",
			"+27 083 653 22699",
			"27 08365322699",
			"27 083653226",
			"27 08a 653 2269",
			"+27 083 6b3 2269",
		}

		za, ok := region.IsoRegionMap["ZA"]
		if !ok {
			t.Errorf("mobile.IsoCountyMap[] = _, %t; want = true", ok)
		}

		for _, tt := range testCases {
			t.Run(tt, func(t *testing.T) {
				_, err := za.SanitiseMobileNumber(tt)
				if err == nil {
					t.Errorf("mobile.Sanitise() = _, %s; want = nil", err)
				}
			})
		}
	})
}
