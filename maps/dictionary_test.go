package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dic := map[string]string{"test": "text context"}

	got := Search(dic, "test")
	want := "text context"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
