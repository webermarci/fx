package fx

func Map[T any](slice []T, fn func(value T) T) []T {
	results := make([]T, len(slice))
	for i, e := range slice {
		results[i] = fn(e)
	}
	return results
}
