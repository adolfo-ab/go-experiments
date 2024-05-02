package main

import "sync"

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			l := getLoggerInstance()
			l.SetLogLevel(i)
			l.Log("Hi mum!")
		}(i)
	}
	wg.Wait()
}
