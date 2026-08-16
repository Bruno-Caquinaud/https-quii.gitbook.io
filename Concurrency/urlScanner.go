package concurrency

type WebsiteChecker func(string) bool
type Result struct {
	url   string
	value bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan Result, len(urls))

	for _, url := range urls {
		go func() {
			resultChannel <- Result{url, wc(url)}
		}()
	}

	for range urls {
		result := <-resultChannel
		results[result.url] = result.value
	}

	return results
}
