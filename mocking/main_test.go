package main

import (
	"bytes"
	"reflect"
	"testing"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)

	return
}

const write = "write"
const sleep = "sleep"

func TestCounting(t *testing.T)  {
	t.Run("Count down 3 to Go", func(t *testing.T) {
		b := &bytes.Buffer{}
		spy := &CountdownOperationsSpy{}

		CountDown(b, spy)

		g := b.String()
		w := `3
2
1
Go!`

		if g != w {
			t.Errorf("got %q, want %q", g, w)
		}
	})
	
	t.Run("test sleep called before every print", func(t *testing.T) {
		spy := &CountdownOperationsSpy{}
		w := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		CountDown(spy, spy)

		if !reflect.DeepEqual(w, spy.Calls) {
			t.Errorf("want %v, got %v", w, spy.Calls)
		}
	})
}
