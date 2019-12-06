package Pages

import (
	"encoding/json"
	"github.com/sharifanani/goWebApp/database/models"
	"github.com/sharifanani/goWebApp/database/queries"
	"github.com/sharifanani/goWebApp/database/responses"
	"io/ioutil"
	"net/http"
	"strings"
)

func closeRequest(r *http.Request) {
	_ = r.Close
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer closeRequest(r)
	var myUser models.User
	_ = json.Unmarshal(body, &myUser)
	if len(myUser.Password) < 6 {
		w.WriteHeader(http.StatusConflict)
	}
	_, err := queries.AddUser(&myUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			w.WriteHeader(http.StatusConflict)
			resp, _ := json.Marshal(
				responses.ServerMsg{Message: "Another user with the same email or username already exists"})
			_, _ = w.Write(resp)
		} else {
			_, _ = w.Write([]byte(err.Error()))
		}
	}
}

func ResetUserTable(w http.ResponseWriter, r *http.Request) {
	_, _ = w, r
	queries.ResetUserTable()
}
