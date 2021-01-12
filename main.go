package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	log.SetFlags(0)
}

// main is a main entry point of this app.
//
// REFERENCES:
//   - https://github.com/mattn/go-sqlite3
func main() {
	db, err := sql.Open("sqlite3", "file:test?mode=memory")
	if err != nil {
		log.Printf("Error: sql.Open (%s)", err)
		os.Exit(1)
	}
	defer db.Close()

	var (
		query string
	)

	query = `
		CREATE TABLE t1 (c1 TEXT)
	`

	if _, err := db.Exec(query); err != nil {
		log.Printf("Error: db.Exec (%s) (%s)", query, err)
		return
	}

	query = `
		INSERT INTO t1 VALUES ('hello')
	`

	if _, err := db.Exec(query); err != nil {
		log.Printf("Error: db.Exec (%s) (%s)", query, err)
		return
	}

	query = `
		SELECT * FROM t1
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error: db.Exec (%s) (%s)", query, err)
		return		
	}
	defer rows.Close()

	for rows.Next() {
		var c1 string
		if err := rows.Scan(&c1); err != nil {
			log.Printf("Error: rows.Scan (%s)", err)
			break
		}

		log.Printf("c1: %v", c1)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error: rows.Err() (%s)", err)
		return
	}
}
