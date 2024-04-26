package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"time"
)

type Activity struct {
	User        string    `json:"user"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}

func processActivity(r io.Reader) error {
	var act Activity

	lr := io.LimitReader(r, 10*1024)
	dec := json.NewDecoder(lr)
	if err := dec.Decode(&act); err != nil {
		return err
	}

	if act.StartTime.After(act.EndTime) {
		return errors.New("start time must be before end time")
	}
	log.Printf("activity: %#v", act)

	return nil
}

func main() {
	if err := processActivity(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
