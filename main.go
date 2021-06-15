package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", listTodosHandler).Methods("GET")
	r.HandleFunc("/add", createTodosHandler).Methods("POST")
	r.HandleFunc("/delete", deleteTodosHandler).Methods("DELETE")
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func listTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"test" : true}`)
	fmt.Println("listTodosHandler")
}

func createTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("createTodosHandler")
}

func deleteTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("deleteTodosHandler")
}
