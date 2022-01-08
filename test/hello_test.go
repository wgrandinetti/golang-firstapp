package main

import "testing"

func TestHello(t *testing.T) {
	t.Helper()
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello2("Chris", "")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello2("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
