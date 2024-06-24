package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("In /hello")
	})

	mux.HandleFunc("/note", checkAuth(processNote))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error happenned: ", err.Error())
		return
	}
}

// Do not use global vars.
var note = Note{
	Value:     "Some initial val",
	UpdatedAt: time.Now(),
}

func processNote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNote(w, r)
	case http.MethodPut:
		putNote(w, r)
	}
}

type Note struct {
	Value     string
	UpdatedAt time.Time
}

func getNote(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(note)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func putNote(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	note.UpdatedAt = time.Now()

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type User struct {
	Username string
	Password string
}

var singleUser = User{
	Username: "ollevche",
	Password: "114477",
}

func checkAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if singleUser.Username != username || singleUser.Password != password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}