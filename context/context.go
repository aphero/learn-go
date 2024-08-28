package context

import (
	"context"
	"fmt"
	"net/http"
)

// The Fetch() and Cancel() functions are only defined in the
// SpyStore struct we're using in place of this interface
type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}

// Version 1
// func LegacyServer(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Our context is tied to the request
// 		ctx := r.Context()

// 		// Create a channel to store our response data in
// 		data := make(chan string, 1)

// 		// Use a goroutine to fetch our response
// 		go func() {
// 			// Send our data to our data channel
// 			data <- store.Fetch()
// 		}()

// 		// This select creates a sort of data race between the response of
// 		// data or the cancellation of the request by the context channel
// 		select {
// 		// Either we retreive our response data from our data channel
// 		case d := <-data:
// 			// and print it
// 			fmt.Fprint(w, d)
// 		// Or we see that our context has been closed (this time via our call to Cancel in a test)
// 		case <-ctx.Done():
// 			// and Cancel the request
// 			store.Cancel()
// 		}
// 	}
// }
