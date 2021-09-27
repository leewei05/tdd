package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	input *bufio.Scanner
}

func NewCLI(store PlayerStore, input io.Reader) *CLI {
	return &CLI{
		store: store,
		input: bufio.NewScanner(input),
	}
}

func (c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.store.RecordWin(extractWinner(userInput))
}

func (c *CLI) readLine() string {
	c.input.Scan()
	return c.input.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
