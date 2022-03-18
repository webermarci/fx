package fx

var table = []struct {
	input []int
}{
	{input: generate(10)},
	{input: generate(100)},
	{input: generate(1000)},
}

func generate(n int) []int {
	var r []int
	for i := 0; i < n; i++ {
		r = append(r, i)
	}
	return r
}
