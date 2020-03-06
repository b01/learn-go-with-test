package main

import (
	"bytes"
	"testing"
)

func TestCounting(t *testing.T)  {
	t.Run("Count down 3", func(t *testing.T) {
		b := bytes.Buffer{}

		CountDown(&b)

		g := b.String()
		w := "3"

		if g != w {
			t.Errorf("got %q, want %q", g, w)
		}
	})
}
