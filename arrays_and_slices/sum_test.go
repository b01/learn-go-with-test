package main

import(
	"testing"
)

func TestSum(t *testing.T)  {
	t.Run("test sum with array", func(t *testing.T) {
		n := []int{1, 2, 3, 4, 5}
		g := Sum(n)
		e := 15

		if g != e {
			t.Errorf("expecting %d got %d given, %v", e, g, n)
		}
	})
}