package region

import (
	"errors"
	"strings"
)

var (
	ErrInvalidNumber = errors.New("invalid number")
)

func StripSpaces(r Region, mobileNumber string) (string, error) {
	out := strings.ReplaceAll(mobileNumber, " ", "")
	return out, nil
}

func StripCountryCode(r Region, mobileNumber string) (string, error) {
	prefixPlus := "+" + r.CC

	if strings.HasPrefix(mobileNumber, r.CC) {
		m := strings.TrimPrefix(mobileNumber, r.CC)
		return m, nil
	} else if strings.HasPrefix(mobileNumber, prefixPlus) {
		m := strings.TrimPrefix(mobileNumber, prefixPlus)
		return m, nil
	} else {
		//do nothing
		return mobileNumber, nil
	}
}

func StripLeadingZero(r Region, mobileNumber string) (string, error) {
	if strings.HasPrefix(mobileNumber, "0") {
		out := strings.TrimPrefix(mobileNumber, "0")
		return out, nil
	}

	return mobileNumber, nil
}

func ValidateAfterCC(r Region, mobileNumber string) (string, error) {
	if len(mobileNumber) != r.lengthAfterCC {
		return "", ErrInvalidNumber
	}

	for _, ch := range mobileNumber {
		if !(ch >= '0' && ch <= '9') {
			return "", ErrInvalidNumber
		}
	}

	return mobileNumber, nil
}

func PrependCC(r Region, mobileNumber string) (string, error) {
	return "+" + r.CC + mobileNumber, nil
}
