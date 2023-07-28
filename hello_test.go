package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("chris")
	want := "hello,chris"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
