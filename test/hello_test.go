package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello3(t *testing.T) {
	got := Hello2("Chris")
	want := "Hello, TestChris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
