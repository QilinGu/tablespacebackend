package main

import (
    "log"
    "net/http"
    "os"
    "database/sql"
    "fmt"
    "strconv"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

var (
    db     *sql.DB = nil
    errd error;
)

func main() {

    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    //Connect to db
	connectToDb();

	//Parse http
    router := gin.New()
    router.Use(gin.Logger())
    router.Static("/static", "static")

    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, string([]byte("You've reached the root directory of Tablespace!")))
       
    })

    router.GET("/welcome/:retrievedname", func(c *gin.Context) {
    	name := c.Param("retrievedname")
        c.String(http.StatusOK, "Hello %s\n", name)
        c.String(http.StatusOK, string([]byte("You've reached the hello directory of Tablespace!")))
       
    })

    router.GET("/menus/:restaurantid", getMenu)

    router.Run(":" + port)

	 //If http request is for menu data
	 	//Get menu ids associated with restaurant
	 		//Get menus associated with previous menu ids
	 			//Get food items associated with current menu id

	 //Generate array of menus
	 //Convert array into json
	 //Return json array

 }

func connectToDb(){

	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL") + "?sslmode=disable")
    if errd != nil {
        log.Fatalf("Error opening database: %q", errd)
    }

}

func getMenu(c *gin.Context) {

	//Gets restaurant id from parameter in path ("<servername>/menus/restaurantid")
	restaurantid,err := strconv.ParseInt(c.Param("restaurantid"), 0, 64)
	if err != nil{
		c.String(http.StatusInternalServerError,
            fmt.Sprintf("Error with restaurant id input: %q", err))
        return
	}
	rows, err := db.Query("SELECT * FROM restaurant WHERE id = $1", restaurantid)
    if err != nil {
        c.String(http.StatusInternalServerError,
            fmt.Sprintf("Error reading restaurant: %q", err))
        return
    }

    defer rows.Close()
    for rows.Next() {
        var row string
        if err := rows.Scan(&row); err != nil {
          c.String(http.StatusInternalServerError,
            fmt.Sprintf("Error scanning menus: %q", err))
            return
        }
        c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", row))
    }
}