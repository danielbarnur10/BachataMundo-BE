package main

import (
	"encoding/json"
	"net/http"
)

// User of Daniel
type User struct {
	username string
	password string
}

var users = [...]User{
	User{
		username: "nitzan",
		password: "1234",
	},
	User{
		username: "daniel",
		password: "12345",
	},
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(users)
	// js, err := json.Marshal(users)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	// w.Write(js)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var username = r.URL.Query().Get("username")
	var foundUsername User

	for i := 0; i < len(users); i++ {
		if users[i].username == username {
			foundUsername = users[i]
		}
	}

	json.NewEncoder(w).Encode(foundUsername)
}
