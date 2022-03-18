package fx

import (
	"fmt"
	"testing"
)

func TestMapInt(t *testing.T) {
	r := Map([]int{1, 2, 3, 4}, func(v int) int {
		return v * 2
	})
	if r[0] != 2 {
		t.Fatal("incorrect calculation")
	}
}

func TestMapString(t *testing.T) {
	r := Map([]string{"hello", "hi", "hey"}, func(v string) string {
		return fmt.Sprintf("%s_test", v)
	})
	if r[0] != "hello_test" {
		t.Fatal("incorrect concat")
	}
}

func BenchmarkMap(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Map(v.input, func(v int) int {
					return v * 2
				})
			}
		})
	}
}
