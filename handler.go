package main

import (
	"io"
	"net/http"
)

type HandlerFunction func(w http.ResponseWriter, r *http.Request)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func handleCharacter(w http.ResponseWriter, r *http.Request) {
	var call HandlerFunction

	switch r.Method {
	case "POST":
		call = addCharacter
	case "GET":
		call = retrieveCharacter
	case "PUT":
		call = updateCharacter
	case "DELETE":
		call = removeCharacter
	}

	call(w, r)
}

func addCharacter(w http.ResponseWriter, r *http.Request) {
	// is NPC flag should be enough
}

func updateCharacter(w http.ResponseWriter, r *http.Request) {

}

func removeCharacter(w http.ResponseWriter, r *http.Request) {
	r.Body.
}

func addNPCCharacter(w http.ResponseWriter, r *http.Request) {

}

func retrieveCharacter(w http.ResponseWriter, r *http.Request) {

}


func startServer() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/character", handleCharacter)

	http.ListenAndServe(":4545", nil)
}

