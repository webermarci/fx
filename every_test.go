package fx

import (
	"fmt"
	"strings"
	"testing"
)

func TestEveryInt(t *testing.T) {
	f := Every([]int{1, 2, 3, 4}, func(v int) bool {
		return v < 5
	})
	if !f {
		t.Fatal("not found")
	}
}

func TestEveryString(t *testing.T) {
	f := Every([]string{"hello", "hi", "hey"}, func(v string) bool {
		return strings.HasPrefix(v, "h")
	})
	if !f {
		t.Fatal("not found")
	}
}

func BenchmarkEvery(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Every(v.input, func(v int) bool {
					return v < i
				})
			}
		})
	}
}
