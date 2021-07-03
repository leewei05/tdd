package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "text context"}

	got := dic.Search("test")
	want := "text context"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
