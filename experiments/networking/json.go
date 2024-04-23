package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Planet struct {
	Name   string  `json:"name"`
	Radius float64 `json:"radius"`
	Mass   float64 `json:"mass"`
}

var data = `
{
	"name": "Mercury",
	"radius": 2439.7,
	"mass": 0.330
}
`

func main() {
	// Create request
	rdr := strings.NewReader(data)

	dec := json.NewDecoder(rdr)

	var pl Planet
	if err := dec.Decode(&pl); err != nil {
		log.Fatal("error: can't decode - %s", err)
	}

	fmt.Printf("got: %+v\n", pl)

	// Create response
	newPlanet := map[string]interface{}{
		"name":   "Jupiter",
		"radius": 3.14,
		"mass":   18.9,
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(newPlanet); err != nil {
		log.Fatal("error: can't encode - %s", err)
	}
}
