package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*
Because this struct is implementing a Fetch() and Cancel() function
it will work to replace the interface we're using for Server(), which is Store

NOTE: Though I'm not entirely sure how the compiler knows that SpyStore and Store are
equivalent due to that fact.
*/
type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		// For each response we get per goroutine
		for _, c := range s.response {
			// If the context is cancelled, trigger that, otherwise default to reading a response
			select {
			// If we see that the context has been cancelled
			case <-ctx.Done():
				// Print that it was cancelled and return early
				log.Println("spy store got cancelled")
				return
			default:
				// Otherwise wait 10 milliseconds and add the response data to our result
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		// Send the result (if any) to our data channel
		data <- result
	}()

	select {
	// Return early with nothing and send a context error in case it hasn't actually been cancelled yet
	// Using this in a select also gives us a way to not halt the function like it normally does when
	// waiting to receive on a channel
	case <-ctx.Done():
		return "", ctx.Err()
		// Otherwise retrieve our data from the channel for our response
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})

	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})
}

// Version 1 (Legacy)
// type LegacySpyStore struct {
// 	response  string
// 	cancelled bool
// 	// We're adding our test suite to the structure so we can call test from it
// 	t *testing.T
// }

// // Fetch a mocked response
// func (s *LegacySpyStore) Fetch() string {
// 	// Fake delay to simulate response time,
// 	// which we'll set higher than the cancellation timeout set in the context
// 	time.Sleep(100 * time.Millisecond)
// 	// For this mock we're just responding with the response we set in the mock struct
// 	return s.response
// }

// // Cancel a mocked request using context
// func (s *LegacySpyStore) Cancel() {
// 	// which just sets cancelled to true for the purpose of our tests
// 	s.cancelled = true
// }

// func (s *LegacySpyStore) assertWasCancelled() {
// 	// Here's why we added the test package to SpyStore
// 	// If we didn't we wouldn't be able to mark this method (not a function) as a helper for tests
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }

// func (s *LegacySpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }

// func TestServer(t *testing.T) {
// 	// data := "hello, world"
// 	// t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
// 	// 	// The data we expect from the SpyStore response

// 	// 	// SpyStore that contains our response
// 	// 	store := &SpyStore{response: data, t: t}

// 	// 	// Create the server using our mock SpyStore
// 	// 	svr := LegacyServer(store)

// 	// 	// Create a new HTTP request, setting method, target and body (which is nil here)
// 	// 	// but using httptest instead of the standard http.Request struct
// 	// 	request := httptest.NewRequest(http.MethodGet, "/", nil)

// 	// 	// Here we're setting the GET request as the parent of the context using WithCancel()
// 	// 	// There are several other conditions you can set with context, including timeout, value triggers, etc.
// 	// 	cancellingCtx, cancel := context.WithCancel(request.Context())

// 	// 	// Using the time package, we wait for 5 milliseconds and then call the cancel function on the context
// 	// 	time.AfterFunc(5*time.Millisecond, cancel)

// 	// 	// Tell the parent context we're cancelling the request
// 	// 	request = request.WithContext(cancellingCtx)

// 	// 	// Create a new recorder to pass to the writer in the Server function
// 	// 	response := httptest.NewRecorder()

// 	// 	// Call ServeHTTP() using our mock response and context request,
// 	// 	// which due to the way we've set up our context will cancel before the request is fulfilled
// 	// 	svr.ServeHTTP(response, request)

// 	// 	// Send an error if we haven't properly cancelled the request
// 	// 	store.assertWasCancelled()
// 	// })

// 	t.Run("returns data from store", func(t *testing.T) {
// 		data := "hello, world"
// 		store := &SpyStore{response: data, t: t}
// 		svr := Server(store)

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		response := httptest.NewRecorder()

// 		svr.ServeHTTP(response, request)

// 		if response.Body.String() != data {
// 			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
// 		}
// 	})
// }
