package main

import(
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test sum with array", func(t *testing.T) {
		n := []int{1, 2, 3, 4, 5}
		g := Sum(n)
		e := 15

		if g != e {
			t.Errorf("expecting %d got %d given, %v", e, g, n)
		}
	})
}
func TestSumAll(t *testing.T) {
	t.Run("Sum all", func(t *testing.T) {
		g := SumAll([]int{1, 3}, []int{2, 3})
		e := []int{4, 5}

		if !reflect.DeepEqual(g, e) {
			t.Errorf("expecting %v got %v", e, g)
		}
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3}, []int{4, 6, 2, 3})
	}
}