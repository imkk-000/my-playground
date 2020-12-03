package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("Failed to %s: %s", msg, err.Error())
	}
}

func main() {
	conn, err := sql.Open("sqlite3", "file:sqlite.db")
	failOnError(err, "create db")

	stmt, err := conn.Prepare("SELECT created_time FROM timer")
	failOnError(err, "prepare statement")
	defer stmt.Close()

	rows, err := stmt.Query()
	failOnError(err, "query row")
	defer rows.Close()

	for rows.Next() {
		var createdTime *time.Time
		err := rows.Scan(&createdTime)
		failOnError(err, "scan row")
		log.Println(createdTime)
	}
}
