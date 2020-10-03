package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	r := mux.NewRouter()
	//http.HandleFunc("user/{id}/set", setUser)
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/user", getAllUsers).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {

	handleRequests()
}
