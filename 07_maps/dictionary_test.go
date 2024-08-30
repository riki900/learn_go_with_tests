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

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionary{}
		word := ""
		definition := ""

		err := dictionary.Add(word, definition)

		AssertErrors(t, err, nil)
		AssertDefintion(t, dictionary, word, definition)

	})

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		AssertErrors(t, err, ErrWordExists)

	})

}

func AssertDefintion(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("error during lookup", err)
	}

	AssertStrings(t, got, definition)

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
