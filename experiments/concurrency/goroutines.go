package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func returnTypeChannel(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s", url, ctype)

}

func siteSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

func siteConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func siteConcurrentChannels(urls []string) {
	ch := make(chan string)
	for _, url := range urls {
		go returnTypeChannel(url, ch)
	}

	for range urls {
		out := <-ch
		fmt.Println(out)
	}

}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	fmt.Println("Serial version:")
	start := time.Now()
	siteSerial(urls)
	fmt.Println(time.Since(start))

	fmt.Println("Concurrent version:")
	start = time.Now()
	siteConcurrent(urls)
	fmt.Println(time.Since(start))

	fmt.Println("Concurrent version using channels:")
	start = time.Now()
	siteConcurrentChannels(urls)
	fmt.Println(time.Since(start))

}
