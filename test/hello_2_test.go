package main

import "testing"

func TestHello2(t *testing.T) {
	got := Hello2("Chris")
	want := "Hello, Pepe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
