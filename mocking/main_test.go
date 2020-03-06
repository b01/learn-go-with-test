package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCounting(t *testing.T)  {
	t.Run("Count down 3", func(t *testing.T) {
		b := &bytes.Buffer{}
		spy := &SpySleeper{}

		CountDown(b, spy)

		g := b.String()
		w := `3
2
1
Go!`

		if g != w {
			t.Errorf("got %q, want %q", g, w)
		}

		if spy.Calls != 4 {
			t.Errorf("got %v, want 4", spy.Calls)
		}
	})
}
