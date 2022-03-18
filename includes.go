package fx

func Includes[T any](slice []T, fn func(value T) bool) bool {
	for _, e := range slice {
		if fn(e) {
			return true
		}
	}
	return false
}
