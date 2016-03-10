package main

import (
    "log"
    "net/http"
    "os"
    "database/sql"
    _ "github.com/lib/pq"
)

var (
    db     *sql.DB = nil
)

func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/database",startDatabase)
}

func startDatabase(w http.ResponseWriter, r *http.Request){
	var err error
    var errd error
	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if errd != nil {
        log.Fatalf("Error opening database: %q", errd)
    }
    w.Write([]byte("Opened database"))
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}