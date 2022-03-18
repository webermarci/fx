package fx

import "golang.org/x/exp/constraints"

func IndexOf[T constraints.Ordered](slice []T, value T) int {
	for i, e := range slice {
		if e == value {
			return i
		}
	}
	return -1
}
