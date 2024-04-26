package main

import (
	"encoding/json"
	"fmt"
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

func (a *Activity) Validate() error {
	if len(a.User) == 0 {
		return fmt.Errorf("invalid user, length must be greater than 0")
	}

	if a.StartTime.After(a.EndTime) {
		return fmt.Errorf("start time must be before end time")
	}

	return nil

}

func processActivity(r io.Reader) error {
	var act Activity
	maxSize := int64(10 * 1024)

	lr := io.LimitReader(r, maxSize)
	dec := json.NewDecoder(lr)
	if err := dec.Decode(&act); err != nil {
		return err
	}

	if err := act.Validate(); err != nil {
		return err
	}

	log.Printf("activity: %#v", act)
	return nil
}

func main() {
	if err := processActivity(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
