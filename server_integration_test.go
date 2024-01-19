package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabeno/poker/v1"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := poker.CreateTempFile(t, `[]`)
	defer cleanDatabase()
	store, err := poker.NewFileSystemStore(database)

	poker.AssertNoError(t, err)

	server := poker.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewGetScoreRequest(player))

		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewLeagueRequest())

		poker.AssertStatus(t, response.Code, http.StatusOK)

		wantedLeague := []poker.Player{
			{"Pepper", 3},
		}
		got := poker.GetLeagueFromReponse(t, response.Body)
		poker.AssertLeague(t, got, wantedLeague)
	})
}
