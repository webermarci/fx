package fx

func Every[T any](slice []T, fn func(value T) bool) bool {
	count := 0
	for _, e := range slice {
		if fn(e) {
			count++
		}
	}
	return count >= len(slice)
}
