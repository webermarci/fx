package fx

import (
	"fmt"
	"log"
	"testing"
)

func TestFilterInt(t *testing.T) {
	r := Filter([]int{1, 2, 3, 4}, func(v int) bool {
		return v > 2
	})
	if len(r) != 2 {
		log.Fatal("not two items found")
	}
}

func TestFilterString(t *testing.T) {
	r := Filter([]string{"hello", "hi", "hey"}, func(v string) bool {
		return len(v) <= 3
	})
	if len(r) != 2 {
		log.Fatal("not two items found")
	}
}

func BenchmarkFilter(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Filter(v.input, func(v int) bool {
					return v < i
				})
			}
		})
	}
}
