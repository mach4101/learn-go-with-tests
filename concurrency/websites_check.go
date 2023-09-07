package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// wg := sync.WaitGroup{}
	// wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			results[url] = wc(url)
			// wg.Done()
		}(url)
	}
	// wg.Wait()
	time.Sleep(2 * time.Second)
	return results
}
