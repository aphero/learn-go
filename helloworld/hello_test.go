package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to humans", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in French", func(t *testing.T) {
		got := Hello("Anna", "French")
		want := "Bonjour, Anna"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in German", func(t *testing.T) {
		got := Hello("Brumhild", "German")
		want := "Hallo, Brumhild"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
