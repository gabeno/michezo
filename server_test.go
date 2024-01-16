package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Bob":   20,
			"Bryce": 10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns score for Bob", func(t *testing.T) {
		request := newGetScoreRequest("Bob")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns score for player Bryce", func(t *testing.T) {
		request := newGetScoreRequest("Bryce")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")

	})

	t.Run("missing player returns 404", func(t *testing.T) {
		request := newGetScoreRequest("Cumilo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)

	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
