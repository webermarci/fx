package fx

import (
	"fmt"
	"testing"
)

func TestFindInt(t *testing.T) {
	_, f := Find([]int{1, 2, 3, 4}, func(v int) bool {
		return v == 2
	})
	if !f {
		t.Fatal("not found")
	}
}

func TestFindString(t *testing.T) {
	_, f := Find([]string{"hello", "hi", "hey"}, func(v string) bool {
		return v == "hi"
	})
	if !f {
		t.Fatal("not found")
	}
}

func TestFindNotFound(t *testing.T) {
	_, f := Find([]int{}, func(v int) bool {
		return v == 2
	})
	if f {
		t.Fatal("found")
	}
}

func BenchmarkFind(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Find(v.input, func(v int) bool {
					return v == i
				})
			}
		})
	}
}
