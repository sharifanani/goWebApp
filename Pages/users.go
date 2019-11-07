package Pages

import (
	"encoding/json"
	"fmt"
	"goWebApp/database/queries"
	"goWebApp/database/tables"
	"io/ioutil"
	"net/http"
)

func closeRequest(r *http.Request) {
	_ = r.Close
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer closeRequest(r)
	var myUser tables.User
	_ = json.Unmarshal(body, &myUser)
	rows, err := queries.AddUser(&myUser)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}
	fmt.Println(rows)
	fmt.Println(err)
	fmt.Println(myUser)
}
