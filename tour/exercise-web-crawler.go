// In this exercise you'll use Go's concurrency features to parallelize a web crawler.
// Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
// Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan [2]string) {
	var mux sync.Mutex
	var wg sync.WaitGroup
	seenUrls := make(map[string]bool)
	firstTime := true
	CrawlRoutine(url, depth, fetcher, seenUrls, ch, &mux, &wg, &firstTime)
	wg.Wait()
	close(ch)
}

// CrawlRoutine performs the actual crawling
func CrawlRoutine(url string, depth int, fetcher Fetcher, seenUrls map[string]bool, ch chan [2]string, mux *sync.Mutex, wg *sync.WaitGroup, firstTime *bool) {
	// Avoids calling wg.Done() the first time CrawlRoutine is called (before recursion)
	if *firstTime {
		*firstTime = false
	} else {
		defer wg.Done()
	}

	// Ends fetching at the set depth
	if depth <= 0 {
		return
	}

	// locks seenUrls to disallow simultaneous access by goroutines
	mux.Lock()
	seenUrls[url] = true
	mux.Unlock()

	// fetch url
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch <- [2]string{url, body}

	// fetches all unseen urls
	for _, u := range urls {
		_, ok := seenUrls[u]
		if !ok {
			wg.Add(1)
			go CrawlRoutine(u, depth-1, fetcher, seenUrls, ch, mux, wg, firstTime)
		}
	}
	return
}

func main() {
	ch := make(chan [2]string)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	for urlAndBody := range ch {
		fmt.Printf("found: %s %q\n", urlAndBody[0], urlAndBody[1])
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
