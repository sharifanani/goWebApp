package queries

import (
	"database/sql"
	"goWebApp/database/tables"
)

func AddUser(user *tables.User) (*sql.Rows, error) {
	db := getDb()
	rows, err := db.Query("INSERT INTO public.users (username, pwd, email) VALUES ($1, $2, $3)",
		user.Username, user.Password, user.Email)
	return rows, err
}
