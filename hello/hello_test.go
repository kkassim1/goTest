package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to ppl", func(t *testing.T) {

		got := Hello("chris", "english")
		want := "hello,chris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("eddie", "spanish")
		want := "hola,eddie"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in french", func(t *testing.T) {

		got := Hello("eddie", "french")
		want := "bonjour,eddie"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in yoruba", func(t *testing.T) {

		got := Hello("eddie", "yoruba")
		want := "báwo,eddie"

		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {

	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
