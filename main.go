package main

import (
    "log"
    "os"
    "database/sql"
    _ "github.com/lib/pq"
)

var (
    db     *sql.DB = nil
)

func main() {
    startDatabase()
}

func startDatabase(){
    var errd error
	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if errd != nil {
        log.Fatalf("Error opening database: %q", errd)
    }
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}