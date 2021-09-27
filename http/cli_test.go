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

	assertPlayerWin(t, store, "Lee")
}
