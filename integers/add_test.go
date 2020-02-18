package integers

import (
	"fmt"
	"testing"
)

func TestAdding(t *testing.T)  {
	w := 4;
	g := Add(2, 2)

	if w != g {
		t.Errorf("Want %d and got %d", w, g)
	}
}

func ExampleAdd() {
	fmt.Println(Add(3, 4))

	// Output: 7
}