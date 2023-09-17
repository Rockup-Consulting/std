package validate

import (
	"time"
)

func IsOlderThan(age int, dob time.Time, now time.Time) bool {
	if now.Year()-dob.Year() > age {
		return true
	}

	if now.Year()-dob.Year() < age {
		return false
	}

	// now.Year - d.Year == age

	if now.Month() > dob.Month() {
		return true
	}

	if now.Month() < dob.Month() {
		return false
	}

	if now.Day() >= dob.Day() {
		return true
	}

	return false
}
