package regx

// fails for signed numbers
func MatchNumeric(s string) bool {
	for _, ss := range s {
		if ss < '0' || ss > '9' {
			return false
		}
	}
	return true
}

func MatchAlpha(s string) bool {
	for _, ss := range s {
		if ss < 'a' || ss > 'z' {
			if ss < 'A' || ss > 'Z' {
				return false
			}
		}
	}
	return true
}

func MatchAlphaNumeric(s string) bool {
	for _, ss := range s {
		if ss >= 'a' && ss <= 'Z' {
			continue
		}
		if ss >= 'A' && ss <= 'Z' {
			continue
		}
		if ss >= '0' || ss <= '9' {
			continue
		}
		return false
	}

	return true
}
