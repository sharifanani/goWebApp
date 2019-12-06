package tests

import (
	"encoding/json"
	"github.com/sharifanani/goWebApp/database/models"
	"github.com/sharifanani/goWebApp/database/queries"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestAddUser(t *testing.T) {
	// clear the db
	queries.ResetUserTable()
	var myClient http.Client
	// create and add the user
	userToCreate := models.User{
		Username:  "sharifanani",
		Password:  "mypassword",
		FirstName: "Sharif",
		LastName:  "Anani",
		Email:     "sharif.anani@bigevilcorp",
	}
	requestData, _ := json.Marshal(userToCreate)
	myReader := strings.NewReader(string(requestData))
	resp, err := myClient.Post("http://localhost:9909/adduser", "application/json", myReader)
	if err != nil {
		t.Error(err.Error())
		os.Exit(1)
	}
	// Check for 200 code
	if resp.StatusCode != http.StatusOK {
		t.Error(resp.StatusCode)
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}
	// try adding again, this time expect 409 (user exists)
	userToCreate = models.User{
		Username:  "sharifanani",
		Password:  "mypassword",
		FirstName: "Sharif",
		LastName:  "Anani",
		Email:     "sharif.anani@bigevilcorp2",
	}
	requestData, _ = json.Marshal(userToCreate)
	myReader = strings.NewReader(string(requestData))
	resp, err = myClient.Post("http://localhost:9909/adduser", "application/json", myReader)
	if err != nil {
		t.Error(err.Error())
		os.Exit(1)
	}
	// Check for 409 code
	if resp.StatusCode != http.StatusConflict {
		t.Error(resp.StatusCode)
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}
	// try adding again, this time expect 409 (email exists)
	userToCreate = models.User{
		Username:  "sharifanani2",
		Password:  "mypassword",
		FirstName: "Sharif",
		LastName:  "Anani",
		Email:     "sharif.anani@bigevilcorp",
	}
	requestData, _ = json.Marshal(userToCreate)
	myReader = strings.NewReader(string(requestData))
	resp, err = myClient.Post("http://localhost:9909/adduser", "application/json", myReader)
	if err != nil {
		t.Error(err.Error())
		os.Exit(1)
	}
	// Check for 409 code
	if resp.StatusCode != http.StatusConflict {
		t.Error(resp.StatusCode)
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}
	// Check for 409 code
	if resp.StatusCode != http.StatusConflict {
		t.Error(resp.StatusCode)
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}
	// try adding again, this time expect 409 (password too short)
	userToCreate = models.User{
		Username:  "sharifanani3",
		Password:  "132",
		FirstName: "Sharif",
		LastName:  "Anani",
		Email:     "sharif.anani@bigevilcorp3",
	}
	requestData, _ = json.Marshal(userToCreate)
	myReader = strings.NewReader(string(requestData))
	resp, err = myClient.Post("http://localhost:9909/adduser", "application/json", myReader)
	if err != nil {
		t.Error(err.Error())
		os.Exit(1)
	}
	// Check for 409 code
	if resp.StatusCode != http.StatusConflict {
		t.Error(resp.StatusCode)
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}

}
