package fx

func Filter[T any](slice []T, fn func(value T) bool) []T {
	temp := make([]T, len(slice))
	count := 0
	for _, e := range slice {
		if fn(e) {
			temp[count] = e
			count++
		}
	}
	return temp[:count]
}
