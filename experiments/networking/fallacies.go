package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// The fallacies of distributed computing: https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/ip", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	const mb = 1 << 20
	r := io.LimitedReader{resp.Body, mb}
	io.Copy(os.Stdout, &r)
}
