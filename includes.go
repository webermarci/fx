package fx

func Includes[T Ordered](slice []T, value T) bool {
	for _, e := range slice {
		if e == value {
			return true
		}
	}
	return false
}
