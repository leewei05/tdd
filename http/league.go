package poker

import (
	"encoding/json"
	"io"
)

func NewLeague(reader io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(reader).Decode(&league)
	if err != nil {
		return nil, err
	}

	return league, nil
}
