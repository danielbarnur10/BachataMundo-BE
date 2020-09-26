package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/user/{id}", getUser)
	http.HandleFunc("/user", getAllUsers)
	log.Fatal(http.ListenAndServe(":1234", nil))
}

func main() {
	handleRequests()
}
