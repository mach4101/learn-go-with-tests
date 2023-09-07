package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	// wg := sync.WaitGroup{}
	// wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			resultChannel <- result{url, wc(url)}
			// wg.Done()
		}(url)
	}
	// wg.Wait()

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
