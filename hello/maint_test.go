package main

import(
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, Chris"
	got := Hello("Chris")

	if want != got {
		t.Errorf("wanted %q got %q", want, got)
	}
}