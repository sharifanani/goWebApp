package queries

import (
	"database/sql"
	"fmt"
	"github.com/sharifanani/goWebApp/database/models"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ResetUserTable() {
	// TODO: add tests to check it's in development mode/environment
	queryPath := filepath.Join("database", "tables", "users.sql")
	bytes, _ := ioutil.ReadFile(queryPath)
	fmt.Println(os.Getwd())
	myQueries := strings.Split(string(bytes), ";\n")
	db := getDb()
	for _, q := range myQueries {
		_, _ = db.Exec(q + ";")
	}
}

func AddUser(user *models.User) (*sql.Rows, error) {
	db := getDb()
	HashPassword(&user.Password) //replaces in-place!
	rows, err := db.Query("INSERT INTO public.users (username, pwd, email) VALUES ($1, $2, $3)",
		user.Username, user.Password, user.Email)
	return rows, err
}
