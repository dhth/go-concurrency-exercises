//////////////////////////////////////////////////////////////////////
//
// Your task is to change the code to limit the crawler to at most one
// page per second, while maintaining concurrency (in other words,
// Crawl() must be called concurrently)
//
// @hint: you can achieve this by adding 3 lines
//

package main

import (
	"fmt"
	"sync"
	"time"
)

// Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a
// real crawler. It crawls until the maximum depth has reached.
func Crawl(url string, depth int, ping <-chan time.Time, wg *sync.WaitGroup) {
	fmt.Printf("[%s] fetching for url: %s, depth: %d\n", time.Now().Format("04:05"), url, depth)
	defer wg.Done()

	if depth <= 0 {
		return
	}

	<-ping
	body, urls, err := fetcher.Fetch(url)
	fmt.Printf("[%s] url: %s - fetcher returned %d urls\n", time.Now().Format("04:05"), url, len(urls))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("[%s] found: %s %q\n", time.Now().Format("04:05"), url, body)

	wg.Add(len(urls))
	for _, u := range urls {
		// Do not remove the `go` keyword, as Crawl() must be
		// called concurrently
		go Crawl(u, depth-1, ping, wg)
	}
	return
}

func main() {
	var wg sync.WaitGroup

	ping := time.Tick(time.Second)

	wg.Add(1)
	Crawl("http://golang.org/", 4, ping, &wg)
	wg.Wait()
}
