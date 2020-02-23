package iteration

import (
	"fmt"
	"testing"
)

func Testloop(t *testing.T) {
	r := Repeat("a", 6)
	e := "aaaaaa"

	if r != e {
		t.Errorf("expected %q got %q", e, r)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", b.N)
	}
}

func TestRepeat(t *testing.T) {
	r := Repeat("c", 2)
	e := "cc"

	if r != e {
		t.Errorf("expected %q got %q", e, r)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("d", 4))

	// Output: dddd
}