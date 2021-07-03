package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "text context"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "text context"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dic.Search("unknown")

		assertErrors(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dic := Dictionary{}
	word := "new"
	definition := "This is a new word"
	dic.Add(word, definition)

	assertDefinitions(t, dic, word, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinitions(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
