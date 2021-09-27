package poker_test

import (
	"strings"
	"testing"

	poker "tdd/http"
)

func TestCLI(t *testing.T) {
	t.Run("record lee win from user input", func(t *testing.T) {
		in := strings.NewReader("Lee wins\n")
		store := &poker.StubPlayerStore{}
		cli := poker.NewCLI(store, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, store, "Lee")
	})

	t.Run("record yee win from user input", func(t *testing.T) {
		in := strings.NewReader("Yee wins\n")
		store := &poker.StubPlayerStore{}
		cli := poker.NewCLI(store, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, store, "Yee")
	})
}
