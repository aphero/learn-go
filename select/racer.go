package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	// The <- operator forces the program to wait until a value
	// is returned from the "blocked" function call, here it's ping()
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("racerTimeout> timeout while waiting on responses from %s and %s", a, b)
	}
}

// Previously used function to create an http request and record the time before response
// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// We're using a struct{} here instead of something like a bool
// because a struct{} is actually the smallest data type available.
//
// We're not sending anything on the channel we're declaring
// so we may as well use the smallest footprint on memory.
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
