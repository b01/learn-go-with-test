package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T)  {
	t.Run("", func(t *testing.T) {
		b := bytes.Buffer{}
		Greet(&b, "Khalifah")

		g := b.String()
		w := "Hello, Khalifah"

		if g != w {
			t.Errorf("got %q, want %q", g, w)
		}
	})
}
