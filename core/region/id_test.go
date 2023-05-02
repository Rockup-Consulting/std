package region_test

import (
	"errors"

	"testing"
	"time"

	"github.com/Rockup-Consulting/go_std/core/region"
)

func TestZaID(t *testing.T) {

	tests := []struct {
		Name    string
		ID      string
		DOB     time.Time
		WantErr error
	}{
		{
			Name:    "valid ID",
			ID:      "9609095105085",
			DOB:     parseDate(t, "1996-09-09"),
			WantErr: nil,
		},
		{
			Name:    "invalid ID",
			ID:      "9609105105085",
			DOB:     parseDate(t, "1996-09-09"),
			WantErr: region.ErrInvalidID,
		},
		{
			Name:    "date of birth mismatch",
			ID:      "9609095105085",
			DOB:     parseDate(t, "1996-09-10"),
			WantErr: region.ErrIdDobMismatch,
		},
		{
			Name:    "born after 2000",
			ID:      "1809095105089",
			DOB:     parseDate(t, "2018-09-09"),
			WantErr: nil,
		},
		{
			Name:    "older than 100",
			ID:      "1809095105089",
			DOB:     parseDate(t, "1918-09-09"),
			WantErr: nil,
		},
	}

	za, ok := region.IsoRegionMap[region.ZA]
	if !ok {
		t.Fatalf("'ZA' not available on IsoRegionMap")
	}

	now := time.Now()

	zaIDValidationFunc, ok := za.IDValidationFuncs[region.ID_Type_NationalID]
	if !ok {
		t.Fatalf("ID_Type_NationalID not in ZA IDValidationFuncs")
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := zaIDValidationFunc(tt.ID, tt.DOB, now)
			if !errors.Is(err, tt.WantErr) {
				var want string
				if tt.WantErr == nil {
					want = "nil"
				} else {
					want = tt.WantErr.Error()
				}

				t.Errorf("za.ValidateNationalID() = %s; want = %s", err, want)
			}
		})
	}
}

func parseDate(t *testing.T, date string) time.Time {
	t.Helper()
	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		t.Fatal(err.Error())
	}

	return time
}
