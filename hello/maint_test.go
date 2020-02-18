package main

import(
	"testing"
)

func TestHello(t *testing.T) {

	assertHelloOutput := func(t *testing.T, w, g string) {
		t.Helper()
		if w != g {
			t.Errorf("wanted %q got %q", w, g)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Chris"
		got := Hello("Chris", "")
		assertHelloOutput(t, want, got)
	})

	t.Run("empty string defaults to World", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("", "")
		assertHelloOutput(t, want, got)
	})

	t.Run("in Spanish", func(t *testing.T) {
		want := "Hola, Elodie"
		got := Hello("Elodie", "es")
		assertHelloOutput(t, want, got)
	})
}