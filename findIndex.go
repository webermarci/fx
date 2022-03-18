package fx

func FindIndex[T any](slice []T, fn func(value T) bool) (int, bool) {
	for i, e := range slice {
		if fn(e) {
			return i, true
		}
	}
	return -1, false
}
