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
	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}
		word := "new"
		definition := "This is a new word"
		err := dic.Add(word, definition)

		assertErrors(t, err, nil)
		assertDefinitions(t, dic, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "new"
		definition := "This is a new word"
		dic := Dictionary{word: definition}
		err := dic.Add(word, definition)

		assertErrors(t, err, ErrWordExists)
		assertDefinitions(t, dic, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		newDefinition := "This is new"
		dic := Dictionary{word: "This is a word"}

		err := dic.Update(word, newDefinition)

		assertErrors(t, err, nil)
		assertDefinitions(t, dic, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		newDefinition := "This is new"
		dic := Dictionary{}

		err := dic.Update(word, newDefinition)

		assertErrors(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "This is a word"
	dic := Dictionary{word: definition}

	dic.Delete(word)

	_, err := dic.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
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
