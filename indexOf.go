package fx

func IndexOf[T Ordered](slice []T, value T) int {
	for i, e := range slice {
		if e == value {
			return i
		}
	}
	return -1
}
