package queries

import (
	"database/sql"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/blake2b"
	"log"
	"strings"
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

func addSalt(inputString *string) {
	*inputString = strings.Join([]string{*inputString, "1234"}, "")
}

func HashPassword(inputPassword *string) {
	addSalt(inputPassword)
	hashed := blake2b.Sum512([]byte(*inputPassword))
	myBytes := make([]byte, len(hashed))
	for i := 0; i < len(hashed); i++ {
		myBytes[i] = hashed[i]
	}
	myString := string(myBytes)
	*inputPassword = myString
}
