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

type fooditem struct {
    name string      `json:"name"`
    description string `json:"description"`
    price string `json:"price"`
    thumbnail string `json:"thumbnail"`
}

type menuinstance struct {
	name string `json:"name"`
	//fooditems []fooditem `json:"fooditems"`
}

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

    //Test endpoint
    router.GET("/welcome/:retrievedname", func(c *gin.Context) {
    	name := c.Param("retrievedname")
        c.String(http.StatusOK, "Hello %s\n\n", name)
        c.String(http.StatusOK, string([]byte("You've reached the test directory of Tablespace!")))
       
    })

    //If http request is for menu data, call getmenu function
    router.GET("/menus/:restaurantid", getMenu)

    router.Run(":" + port)

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
	//Get menu ids associated with restaurant
	restaurantidrows, err := db.Query("SELECT menuid FROM restaurantmenus WHERE restaurantid = $1", restaurantid)
    if err != nil {
        c.String(http.StatusInternalServerError,
            fmt.Sprintf("Error reading restaurant: %q", err))
        return
    }

    defer restaurantidrows.Close()
    for restaurantidrows.Next() {
        var menuid string
        if err := restaurantidrows.Scan(&menuid); err != nil {
          c.String(http.StatusInternalServerError,
            fmt.Sprintf("Error scanning menus: %q", err))
            return
        }

		c.String(http.StatusOK, fmt.Sprintf("  Menu id read from DB: %s\n", menuid))

		//Start: Get menus associated with previous menu ids
		menuidint, err := strconv.ParseInt(menuid, 0, 64)
		if err != nil{
			c.String(http.StatusInternalServerError,
	            fmt.Sprintf("Error with getting menu id: %q", err))
	        return
		}
        
        menuidrows, err := db.Query("SELECT name FROM menu WHERE id = $1", menuidint)
	    if err != nil {
	        c.String(http.StatusInternalServerError,
	            fmt.Sprintf("Error reading menu: %q", err))
	        return
	    }

	    defer menuidrows.Close()
	    for menuidrows.Next() {
	        var menuname string
	        if err := menuidrows.Scan(&menuname); err != nil {
	          c.String(http.StatusInternalServerError,
	            fmt.Sprintf("Error scanning menus: %q", err))
	            return
	        }

	        c.String(http.StatusOK, fmt.Sprintf("   Menu name read from DB: %s\n", menuname))

	        //START: Get food items associated with current menu id
	        
	        menuitemsrows, err := db.Query("SELECT foodid FROM menuitems WHERE menuid = $1", menuidint)
		    if err != nil {
		        c.String(http.StatusInternalServerError,
		            fmt.Sprintf("Error reading menu: %q", err))
		        return
		    }

		    defer menuitemsrows.Close()
		    for menuitemsrows.Next() {
		        var fooditemid string
		        if err := menuitemsrows.Scan(&fooditemid); err != nil {
		          c.String(http.StatusInternalServerError,
		            fmt.Sprintf("Error scanning for menu items: %q", err))
		            return
		        }
		        
		        c.String(http.StatusOK, fmt.Sprintf("    Food item id read from DB: %s\n", fooditemid))

		        //START: Get food item detais 
		        fooditemidint, err := strconv.ParseInt(fooditemid, 0, 64)
				if err != nil{
					c.String(http.StatusInternalServerError,
			            fmt.Sprintf("Error with getting menu id: %q", err))
			        return
				}
		        
		        fooditemidrows, err := db.Query("SELECT name, description, price, itempicture FROM fooditem WHERE id = $1", fooditemidint)
			    if err != nil {
			        c.String(http.StatusInternalServerError,
			            fmt.Sprintf("Error reading food item name: %q", err))
			        return
			    }

			    defer fooditemidrows.Close()
			    for fooditemidrows.Next() {
			        var fooditemname string
					var fooditemdescription string 
					var fooditemprice string
					var fooditemthumbnail string
			        if err := fooditemidrows.Scan(&fooditemname, &fooditemdescription, &fooditemprice, &fooditemthumbnail); err != nil {
			          c.String(http.StatusInternalServerError,
			            fmt.Sprintf("Error parsing food item: %q", err))
			            return
			        }

			        c.String(http.StatusOK, fmt.Sprintf("     Food item name read from DB: %s\n", fooditemname))
			        c.String(http.StatusOK, fmt.Sprintf("     Food item description read from DB: %s\n", fooditemdescription))
			        c.String(http.StatusOK, fmt.Sprintf("     Food item price read from DB: %s\n", fooditemprice))
			        c.String(http.StatusOK, fmt.Sprintf("     Food item thumbnail read from DB: %s\n", fooditemthumbnail))

			    }
		        //END: Get food item detais

		    }
	        //END: Get food items associated with current menu id
	    } 
	    //END: Get menus associated with previous menu ids
       
    }
}