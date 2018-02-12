package main

import (
	"io"
	"net/http"
	"encoding/json"
)

type HandlerFunction func(w http.ResponseWriter, r *http.Request)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func getEncounter(w http.ResponseWriter, r *http.Request) {
	var request Encounter
	json.Unmarshal(extractBody(w, r), &request)

	createEncounter(request)

	io.WriteString(w, request.getString())
}

func extractBody(w http.ResponseWriter, r *http.Request) []byte {
	var body []byte
	res, err := r.Body.Read(body)
	if (err != nil) {
		io.WriteString(w, "Bad request")
	}

	return body
}

func getNextInLineup(w http.ResponseWriter, r *http.Request) {

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

