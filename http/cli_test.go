package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Lee wins\n")
	store := &StubPlayerStore{}
	cli := &CLI{store, in}
	cli.PlayPoker()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := store.winCalls[0]
	want := "Lee"
	if got != want {
		t.Errorf("didn't record correct winner, got %q, want %q", got, want)
	}
}
