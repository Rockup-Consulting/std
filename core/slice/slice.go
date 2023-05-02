package slice

import "golang.org/x/exp/constraints"

func Contains[T comparable](a []T, b T) bool {
	for _, val := range a {
		if val == b {
			return true
		}
	}

	return false
}

func Max[T constraints.Ordered](a []T) T {
	if len(a) == 0 {
		var zero T
		return zero
	}
	max := a[0]

	for _, val := range a {
		if val > max {
			max = val
		}
	}

	return max
}

func Min[T constraints.Ordered](a []T) T {
	if len(a) == 0 {
		var zero T
		return zero
	}
	min := a[0]

	for _, val := range a {
		if val < min {
			min = val
		}
	}

	return min
}

func RemoveDuplicates[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	list := []T{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range slice {
		_, value := keys[entry]
		if !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Prepend[T comparable](a T, b []T) []T {
	out := make([]T, len(b)+1)
	copy(out[1:], b)

	out[0] = a
	return out
}

func Remove[T comparable](s []T, index int) []T {
	return append(s[0:index], s[index+1:]...)
}
