package fx

import (
	"fmt"
	"testing"
)

func TestMinInt(t *testing.T) {
	r := Min([]int{4, 3, 2, 1})
	if r != 1 {
		t.Fatal("incorrect calculation")
	}
}

func TestMinFloat(t *testing.T) {
	r := Min([]float64{1.5, 2.5, 3.5, 4.5})
	if r != 1.5 {
		t.Fatal("incorrect calculation")
	}
}

func TestMinEmptySlice(t *testing.T) {
	r := Min([]int{})
	if r != 0 {
		t.Fatal("should be 0")
	}
}

func BenchmarkMin(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Min(v.input)
			}
		})
	}
}
