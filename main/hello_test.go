package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "en")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("emtpy string defaults to world", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("hello world in Spanish", func(t *testing.T) {
		got := Hello("", "es")
		want := "Hola mundo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("hello Chris in spanish", func(t *testing.T) {
		got := Hello("Chris", "es")
		want := "Hola Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("hello Chris in french", func(t *testing.T) {
		got := Hello("Chris", "fr")
		want := "Bonjour Chris"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
