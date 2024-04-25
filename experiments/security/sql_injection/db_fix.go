package main

import (
	"database/sql"
	"time"
)

var (
	schemaSQLFix string
	insertSQLFix string
)

func createTableFix(db *sql.DB) error {
	_, err := db.Exec(schemaSQL)
	return err
}

func insertLogFix(db *sql.DB, time time.Time, message string) error {
	_, err := db.Exec(insertSQL, time, message)
	return err
}
