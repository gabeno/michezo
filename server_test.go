package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabeno/poker/v1"
)

func TestGETPlayers(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{
			"Bob":   20,
			"Bryce": 10,
		},
		[]string{},
		nil,
	}
	server := poker.NewPlayerServer(&store)

	t.Run("returns score for Bob", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Bob")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns score for player Bryce", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Bryce")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "10")

	})

	t.Run("missing player returns 404", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Cumilo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusNotFound)

	})
}

func TestStoreWins(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{},
		[]string{},
		nil,
	}
	server := poker.NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request := poker.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.WinCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin, want %d", len(store.WinCalls), 1)
		}

		if store.WinCalls[0] != player {
			t.Errorf("got %s wanted %s", store.WinCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns league table as json", func(t *testing.T) {
		wantedLeague := []poker.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		stubStore := &poker.StubPlayerStore{nil, nil, wantedLeague}
		server := poker.NewPlayerServer(stubStore)

		request := poker.NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := poker.GetLeagueFromReponse(t, response.Body)
		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertLeague(t, got, wantedLeague)
		poker.AssertContentType(t, response, "application/json")
	})
}
