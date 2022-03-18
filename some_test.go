package fx

import (
	"fmt"
	"testing"
)

func TestSomeInt(t *testing.T) {
	f := Some([]int{1, 2, 3, 4}, func(v int) bool {
		return v%2 == 0
	})
	if !f {
		t.Fatal("not found")
	}
}

func TestSomeString(t *testing.T) {
	f := Some([]string{"hello", "hi", "hey"}, func(v string) bool {
		return len(v) <= 3
	})
	if !f {
		t.Fatal("not found")
	}
}

func TestSomeNotFound(t *testing.T) {
	f := Some([]int{1, 2, 3, 4}, func(v int) bool {
		return v%5 == 0
	})
	if f {
		t.Fatal("found")
	}
}

func BenchmarkSome(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Some(v.input, func(v int) bool {
					return v == i
				})
			}
		})
	}
}
