package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	log.SetFlags(0)
}

func open(connString string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, fmt.Errorf("Error: sql.Open (%w)", err)
	}

	return db, err
}

func exec(db *sql.DB, query string) (sql.Result, error) {
	r, err := db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("Error: db.Exec (%w) (%s)", err, query)
	}

	return r, err
}

func query(db *sql.DB, q string) (*sql.Rows, error) {
	r, err := db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("Error: db.Query (%w) (%s)", err, q)
	}

	return r, err
}

// main is a main entry point of this app.
//
// REFERENCES:
//   - https://github.com/mattn/go-sqlite3
//   - https://golang.org/pkg/database/sql/
func main() {
	db, err := open("file:test?mode=memory")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	var (
		q string
	)

	q = `
		CREATE TABLE t1 (id INTEGER, c1 TEXT)
	`

	if _, err := exec(db, q); err != nil {
		log.Println(err)
		return
	}

	for i, v := range []string{"world", "hello"} {
		// FIXME: use placeholder
		q = fmt.Sprintf("INSERT INTO t1 VALUES (%d, '%s')", i+1, v)

		if _, err := exec(db, q); err != nil {
			log.Println(err)
			return
		}
	}

	q = `
		SELECT * FROM t1 ORDER BY id DESC
	`

	rows, err := query(db, q)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id int
			c1 string
		)

		if err := rows.Scan(&id, &c1); err != nil {
			log.Printf("Error: rows.Scan (%s)", err)
			break
		}

		log.Printf("row: %v, %v", id, c1)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error: rows.Err() (%s)", err)
		return
	}
}
