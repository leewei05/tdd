package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

// This method implments PlayerStore interface
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"lee":  20,
			"greg": 10,
		},
	}

	// needs to pass store as a pointer because GetPlayerScore is a pointer receiver
	server := &PlayerServer{&store}

	t.Run("returns Lee's score", func(t *testing.T) {
		request := newPlayerHTTPRequest("lee")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})

	t.Run("returns Greg's score", func(t *testing.T) {
		request := newPlayerHTTPRequest("greg")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
	})
}

func newPlayerHTTPRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
