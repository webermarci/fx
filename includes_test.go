package fx

import (
	"fmt"
	"testing"
)

func TestIncludesInt(t *testing.T) {
	f := Includes([]int{1, 2, 3, 4}, 1)
	if !f {
		t.Fatal("not found")
	}
}

func TestIncludesString(t *testing.T) {
	f := Includes([]string{"hello", "hi", "hey"}, "hi")
	if !f {
		t.Fatal("not found")
	}
}

func TestIncludesNotFound(t *testing.T) {
	f := Includes([]int{}, 1)
	if f {
		t.Fatal("found")
	}
}

func BenchmarkIncludes(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Includes(v.input, i)
			}
		})
	}
}
