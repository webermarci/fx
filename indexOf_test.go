package fx

import (
	"fmt"
	"testing"
)

func TestIndexOfInt(t *testing.T) {
	i := IndexOf([]int{1, 2, 3, 4}, 2)
	if i != 1 {
		t.Fatal("not correct index")
	}
}

func TestIndexOfString(t *testing.T) {
	i := IndexOf([]string{"hello", "hi", "hey"}, "hi")
	if i != 1 {
		t.Fatal("not correct index")
	}
}

func TestIndexOfNotFound(t *testing.T) {
	i := IndexOf([]string{}, "hi")
	if i != -1 {
		t.Fatal("not correct index")
	}
}

func BenchmarkIndexOf(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IndexOf(v.input, i)
			}
		})
	}
}
