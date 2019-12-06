package Pages

import (
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/adduser", AddUser)
	http.HandleFunc("/resetusers", ResetUserTable)
}
