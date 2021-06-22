package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "todosbackend"
)

func main() {
	// set up db connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer closeConnection(db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected")

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

func closeConnection(db *sql.DB) {
	fmt.Println("closing connection...")
	db.Close()
}
