package fx

import (
	"fmt"
	"testing"
)

func TestReduceInt(t *testing.T) {
	r := Reduce([]int{1, 2, 3, 4}, func(previous, current int) int {
		return previous + current
	})
	if r != 10 {
		t.Fatal("incorrect calculation")
	}
}

func TestReduceFloat(t *testing.T) {
	r := Reduce([]float64{1.5, 2.5, 3.5, 4.5}, func(previous, current float64) float64 {
		return previous + current
	})
	if r != 12 {
		t.Fatal("incorrect calculation")
	}
}

func BenchmarkReduce(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Reduce(v.input, func(previous, current int) int {
					return previous + current
				})
			}
		})
	}
}
