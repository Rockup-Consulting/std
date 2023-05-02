package conf

func (m Map) findDuplicateFieldKeys() error {

	testMap := make(map[FieldKeys]string)

	for k, v := range m {
		clashKey, ok := testMap[v.Keys]
		if !ok {
			testMap[v.Keys] = k
		} else {
			return errDuplicateFieldKeys([]string{clashKey, k})
		}
	}

	return nil
}
