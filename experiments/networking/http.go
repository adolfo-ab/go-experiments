package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Travel struct {
	Origin      string  `json: "origin"`
	Destination string  `json: "destination"`
	Price       float64 `json: "price"`
}

func main() {
	// GET request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	//POST request
	travel := &Travel{
		Origin:      "Bergen",
		Destination: "Oslo",
		Price:       329.5,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(travel); err != nil {
		log.Fatalf("error: can't encode travel - %s", err)
	}

	resp, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
