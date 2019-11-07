package queries

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func getDb() *sql.DB {
	connStr := "user=anani dbname=lims port=5050 host=localhost password=mypassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		return db
	}
}
