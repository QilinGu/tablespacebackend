package main

import (
    "bytes"
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"
    "database/sql"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

var (
    db     *sql.DB = nil
    router gin;
)

func main() {

	var err error
    var errd error
    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    //Parse http

    router = gin.New()
    router.Use(gin.Logger())
    router.Static("/static", "static")

    router.GET("/menu", func(c *gin.Context) {
        c.String(http.StatusOK, string([]byte("**hi!**")))
       
    })


	//Connect to db
	connectToDb();


 //If http request is for menu data
 	//Get menu ids associated with restaurant
 		//Get menus associated with previous menu ids
 			//Get food items associated with current menu id

 //Generate array of menus
 //Convert array into json
 //Return json array

 }

func connectToDb(){

	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if errd != nil {
        log.Fatalf("Error opening database: %q", errd)
    }

}