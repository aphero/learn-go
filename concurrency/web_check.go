package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// SEND (<-) a "result" struct to the resultChannel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// RECEIVE (also <-, which is marginally confusing) the values from the resultChannel
		r := <-resultChannel
		// and then record the result in the results map
		results[r.string] = r.bool
	}

	return results
}
