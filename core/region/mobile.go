package region

type SanitiseMobileFunc func(c Region, mobileNumber string) (string, error)

// https://countrycode.org/
func (c Region) SanitiseMobileNumber(mobileNumber string) (string, error) {
	out := mobileNumber
	var err error

	for _, f := range c.sanitiseMobile {
		out, err = f(c, out)
		if err != nil {
			return "", err
		}
	}

	return out, nil
}
