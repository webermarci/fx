package fx

func Min[T Ordered](slice []T) T {
	if len(slice) == 0 {
		return *new(T)
	}

	temp := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < temp {
			temp = slice[i]
		}
	}
	return temp
}
