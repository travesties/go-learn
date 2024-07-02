package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"travis": 20,
			"bob":    10,
		},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("returns Travis' score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/travis", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		// This method of composing an error provides a more concise call.
		assertResponseBody(t, response.Body.String(), "20")

		if err := statusEquals(response.Code, http.StatusOK); err != nil {
			t.Error(err)
		}
	})

	t.Run("returns Bob's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/bob", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		// This method of composing an error has the benefit of ensuring
		// that calls to t.Error or t.Fatal aren't buried in helper methods,
		// which may make it easier to avoid weird debug issues.
		if err := responseBodyEquals(response.Body.String(), "10"); err != nil {
			t.Error(err)
		}

		if err := statusEquals(response.Code, http.StatusOK); err != nil {
			t.Error(err)
		}
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/herpderp", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if err := statusEquals(response.Code, http.StatusNotFound); err != nil {
			t.Error(err)
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}

	server := &PlayerServer{&store}

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "travis"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if err := statusEquals(response.Code, http.StatusAccepted); err != nil {
			t.Error(err)
		}

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf(
				"did not store correct winner: get %q, want %q",
				store.winCalls[0],
				player)
		}
	})

}
func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func responseBodyEquals(got, want string) error {
	if got != want {
		return fmt.Errorf("wrong response: got %q, want %q", got, want)
	}
	return nil
}

func statusEquals(got, want int) error {
	if got != want {
		return fmt.Errorf("wrong status: got %d, want %d", got, want)
	}
	return nil
}

func newPostWinRequest(name string) *http.Request {
	target := fmt.Sprintf("/players/%s", name)
	return httptest.NewRequest(http.MethodPost, target, nil)
}

func newGetScoreRequest(name string) *http.Request {
	target := fmt.Sprintf("/players/%s", name)
	return httptest.NewRequest(http.MethodGet, target, nil)
}
