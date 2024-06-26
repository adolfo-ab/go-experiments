package main

import (
	"database/sql"
	"fmt"
	"time"
)

var (
	schemaSQL string
	insertSQL string
)

func createTable(db *sql.DB) error {
	_, err := db.Exec(schemaSQL)
	return err
}

func insertLog(db *sql.DB, time time.Time, message string) error {
	ts := time.Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf(insertSQL, ts, message)
	_, err := db.Exec(sql)
	return err
}
