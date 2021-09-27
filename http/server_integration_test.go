package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDB := createTempFile(t, `[]`)
	defer cleanDB()
	store, err := NewFileSystemPlayerStore(database)
	AssertNoError(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPlayerPOSTRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPlayerPOSTRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPlayerPOSTRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newPlayerGETRequest(player))
		AssertStatus(t, response.Code, http.StatusOK)

		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		AssertStatus(t, response.Code, http.StatusOK)

		got := getLeagueListFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		AssertLeague(t, got, want)
	})
}
