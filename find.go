package fx

func Find[T any](slice []T, fn func(value T) bool) (T, bool) {
	for _, e := range slice {
		if fn(e) {
			return e, true
		}
	}
	return *new(T), false
}
