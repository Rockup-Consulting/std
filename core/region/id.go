package region

import (
	"errors"
	"time"

	"github.com/Rockup-Consulting/go_std/core/regx"
	"github.com/Rockup-Consulting/go_std/core/validate"
)

type IDValidationFunc = func(id string, dob time.Time, now time.Time) error

const (
	ID_Type_NationalID = "National ID"
	ID_Type_Passport   = "Passport"
	ID_Type_Drivers    = "Driver's License"
)

// ====================================================================
// Errors

var ErrIdDobMismatch = errors.New("mismatch: ID and Date of Birth")

var ErrInvalidID = errors.New("invalid ID")

func IsIdDobMismatchErr(err error) bool {
	return errors.Is(err, ErrIdDobMismatch)
}

func IsInvalidIDErr(err error) bool {
	return errors.Is(err, ErrInvalidID)
}

// ====================================================================
// Validation Function

// assumes dob
func validateZaID(id string, dob time.Time, now time.Time) error {
	//	The format is [YYMMDD] [GGGG] [C] [A] [Z]
	// • 	YYMMDD for the ID holder's date of birth
	// • 	[GGGG] for ID holder's gender. 0000 - 4999 for female and 5000 - 9999 for
	// 		male.
	//	• 	[C] is their citizenship status. O is a South African Citizen and 1 is a
	//		permanent resident.
	//	• 	[A] can be any number but, at the time of writing this post, is usually an 8.
	//	• 	[Z] is a check digit which uses the Luhn Algorithm to validate the rest of the
	//		ID number.

	if len(id) != 13 {
		return ErrInvalidID
	}

	ok := validate.LuhnDigit(id)
	if !ok {
		return ErrInvalidID
	}

	// compare to DOB
	idDob := id[:6]
	var t time.Time
	var err error

	if validate.IsOlderThan(100, dob, now) {
		t, err = time.Parse("20060102", "19"+idDob)
	} else {
		t, err = time.Parse("060102", idDob)
	}
	if err != nil {
		return ErrInvalidID
	}

	if t != dob {
		return ErrIdDobMismatch
	}

	return nil
}

func validateZaPassport(passport string, dob time.Time, now time.Time) error {
	if len(passport) != 9 {
		return ErrInvalidID
	}

	if !regx.MatchAlpha(string(passport[0])) {
		return ErrInvalidID
	}

	if !regx.MatchNumeric(passport[1:]) {
		return ErrInvalidID
	}

	return nil
}
