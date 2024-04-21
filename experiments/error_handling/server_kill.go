package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func killServer(pidFile string) error {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return err
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}

	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}

func main() {
	if err := killServer("experiments/error_handling/server.pid"); err != nil {
		log.Fatalf("error: %s", err)
		os.Exit(1)
	}
}
