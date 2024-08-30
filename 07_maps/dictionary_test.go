package main

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	AssertStrings(t, got, want)

}

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
