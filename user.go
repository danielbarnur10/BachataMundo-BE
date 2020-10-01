package main

import (
	"encoding/json"
	"net/http"
)

// User of Daniel
type User struct {
	Username string
	Password string
}

var users = [...]User{
	User{
		Username: "nitzan",
		Password: "1234",
	},
	User{
		Username: "daniel",
		Password: "12345",
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)
	// w.Write(js)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var username = r.URL.Query().Get("username")
	var foundUsername User

	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			foundUsername = users[i]
		}
	}

	json.NewEncoder(w).Encode(foundUsername)
}
