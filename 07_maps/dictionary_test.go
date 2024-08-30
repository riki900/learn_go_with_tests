package main

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("word found", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		AssertStrings(t, got, want)

	})

	t.Run("word not found", func(t *testing.T) {

		_, err := dictionary.Search("not found")

		AssertErrors(t, err, ErrNotFound)

	})

}

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func AssertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
