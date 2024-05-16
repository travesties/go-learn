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
		// We need to give each anonymous function its own url parameter.
		// We can access url in the lexical scope, but each function
		// invocation will be handling the same variable (no good)
		go func(u string) {
			// Writing directly to the map makes our code vulnerable to
			// a race condition where two goroutines try to write to the
			// map at the same time. We can solve this with channels.
			//results[u] = wc(u)

			// We send a result to a channel from within our goroutines,
			// ensuring that no race conditions are possible. A channel
			// is essentially a FIFO queue used concurrent code.
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// With the use of a channel, we can receive its values one-by-one,
		// avoiding the race condition of two go routines trying to write to
		// the same map location.

		// Channels, by default, are blocking until ready to send/receive.
		// In this instance, the channel blocks until the goroutines sync.
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
