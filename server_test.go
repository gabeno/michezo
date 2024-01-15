package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns score for Bob", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Bob", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns score for player Bryce", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Bryce", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
