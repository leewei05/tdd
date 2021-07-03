package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dic := map[string]string{"test": "text context"}

	got := Search(dic, "test")
	want := "text context"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
