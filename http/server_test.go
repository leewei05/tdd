package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}

	store := StubPlayerStore{nil, nil, wantedLeague}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueListFromResponse(t, response.Body)
		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
		AssertContentType(t, response, jsonContentType)
	})
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"lee":  20,
			"greg": 10,
		},
	}
	// needs to pass store as a pointer because GetPlayerScore is a pointer receiver
	server := NewPlayerServer(&store)

	t.Run("returns Lee's score", func(t *testing.T) {
		request := newPlayerGETRequest("lee")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseBody(t, response.Body.String(), "20")
		AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns Greg's score", func(t *testing.T) {
		request := newPlayerGETRequest("greg")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertResponseBody(t, response.Body.String(), "10")
		AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 for non-existing player", func(t *testing.T) {
		request := newPlayerGETRequest("pig")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{},
	}
	server := NewPlayerServer(store)
	player := "lee"

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request := newPlayerPOSTRequest("lee")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)
		AssertPlayerWin(t, store, player)
	})
}

func newPlayerPOSTRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPlayerGETRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func getLeagueListFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return league
}
