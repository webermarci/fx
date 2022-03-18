package fx

import (
	"fmt"
	"testing"
)

func TestMaxInt(t *testing.T) {
	r := Max([]int{1, 2, 3, 4})
	if r != 4 {
		t.Fatal("incorrect calculation")
	}
}

func TestMaxFloat(t *testing.T) {
	r := Max([]float64{1.5, 2.5, 3.5, 4.5})
	if r != 4.5 {
		t.Fatal("incorrect calculation")
	}
}

func TestMaxEmptySlice(t *testing.T) {
	r := Max([]int{})
	if r != 0 {
		t.Fatal("should be 0")
	}
}

func BenchmarkMax(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Max(v.input)
			}
		})
	}
}
