package validate

import "strconv"

// Validates the checkDigit
// Double even numbered digits within ID
// Sum with remaining uneven digits
//
// If this sum % 10 == 0, the ID number is valid.
func LuhnDigit(id string) bool {

	var totalSum int
	for i, v := range id {
		vInt, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err.Error())
		}

		if i%2 == 0 {
			totalSum += vInt
		} else {
			if (vInt * 2) > 9 {
				totalSum += (vInt * 2) - 9
			} else {
				totalSum += vInt * 2
			}
		}
	}

	return totalSum%10 == 0
}
