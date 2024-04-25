package main

import (
	"database/sql"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	msg := string(data)

	db, err := sql.Open("postgres", "user=postgres password=s3cr3t sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := createTable(db); err != nil {
		log.Fatal(err)
	}

	if err := insertLog(db, time.Now(), msg); err != nil {
		log.Fatal(err)
	}
}
