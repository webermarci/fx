package fx

func Reduce[T any](slice []T, fn func(previous T, current T) T) T {
	var result T
	for _, e := range slice {
		result = fn(result, e)
	}
	return result
}
