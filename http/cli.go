package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	input io.Reader
}

func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.input)
	reader.Scan()
	player := extractWinner(reader.Text())
	c.store.RecordWin(player)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
