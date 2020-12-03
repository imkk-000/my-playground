package main

import (
	"database/sql"
	"log"
	"reflect"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("Failed to %s: %s", msg, err.Error())
	}
}

type Entity struct {
	ID          sql.NullInt32 `db:"id"`
	CreatedTime sql.NullTime  `db:"created_time"`
}

func main() {
	conn, err := sqlx.Open("sqlite3", "file:sqlite.db")
	failOnError(err, "create db")

	err = conn.Ping()
	failOnError(err, "ping db")

	stmt, err := conn.Preparex("SELECT * FROM timer")
	failOnError(err, "prepare statement")
	defer stmt.Close()

	rows, err := stmt.Queryx()
	failOnError(err, "query row")
	defer rows.Close()

	for rows.Next() {
		// var entity Entity
		// scannerPointers := []interface{}{
		// 	&entity.ID,
		// 	&entity.CreatedTime,
		// }
		row, err := rows.SliceScan()
		failOnError(err, "scan row")
		log.Printf("row: %+v", row)

		for i, v := range row {
			vType := reflect.TypeOf(v)
			if vType == nil {
				log.Println("null")
				continue
			}
			log.Printf("%d:  type=%T ", i, v)
			switch vType.Kind() {
			case reflect.Int64:
				log.Printf("(int64)")
			case reflect.Struct:
				log.Printf("(struct)")
			}
			log.Println()
		}
		// log.Printf("struct: %+v", entity)
		// log.Println(entity.ID.Value())
		// log.Println(entity.CreatedTime.Value())
		// log.Printf("id=%d, created_time=%v", entity.ID.Int32, entity.CreatedTime.Time)
	}
}
