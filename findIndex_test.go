package fx

import (
	"fmt"
	"testing"
)

func TestFindIndexInt(t *testing.T) {
	i, f := FindIndex([]int{1, 2, 3, 4}, func(v int) bool {
		return v == 2
	})
	if !f {
		t.Fatal("not found")
	}
	if i != 1 {
		t.Fatal("not correct index")
	}
}

func TestFindIndexString(t *testing.T) {
	i, f := FindIndex([]string{"hello", "hi", "hey"}, func(v string) bool {
		return v == "hi"
	})
	if !f {
		t.Fatal("not found")
	}
	if i != 1 {
		t.Fatal("not correct index")
	}
}

func TestFindIndexNotFound(t *testing.T) {
	_, f := FindIndex([]int{}, func(v int) bool {
		return v == 2
	})
	if f {
		t.Fatal("found")
	}
}

func BenchmarkFindIndex(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindIndex(v.input, func(v int) bool {
					return v == i
				})
			}
		})
	}
}
