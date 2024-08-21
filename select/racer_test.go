package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(10 * time.Millisecond)

		// Closes the servers (or whatever function you want to defer) at the end of the function.
		// Useful for making your code cleaner in the case that you are starting a process that should
		// be closed later
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Errorf("did not expect an error but got one %v", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected a timeout error, but didn't get one")
		}
	})
}

// From the guide, here's the breakdown of these servers:
// NewServer
// - creates a local server with an open port
// - uses a handler function to perform an action
// HandlerFunc
// - Uses a ResponseWriter to fake a response
// - Uses a Request to fake a request
// In this case we just need to delay for a time and then respond with a 200 to fake that we got a simple OK
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
