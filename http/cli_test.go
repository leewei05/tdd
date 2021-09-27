package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record lee win from user input", func(t *testing.T) {
		in := strings.NewReader("Lee wins\n")
		store := &StubPlayerStore{}
		cli := &CLI{store, in}
		cli.PlayPoker()

		assertPlayerWin(t, store, "Lee")
	})

	t.Run("record yee win from user input", func(t *testing.T) {
		in := strings.NewReader("Yee wins\n")
		store := &StubPlayerStore{}
		cli := &CLI{store, in}
		cli.PlayPoker()

		assertPlayerWin(t, store, "Yee")
	})
}
