package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := InMemoryPlayerStore{}
	server := PlayerServer{&store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPlayerPOSTRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPlayerPOSTRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPlayerPOSTRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newPlayerGETRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "123")
}
