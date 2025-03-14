package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"FASManagementSystem/internal/database"

)
type album struct {
	Title  string  `json:"title"`
    Artist string  `json:"artist"`
    ID     string  `json:"id"`
    Price  float64 `json:"price"`
}
// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
func main() {
	database.ConnectDB()
	fmt.Println("Hello World")

    router := gin.Default()
    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")
	
}

// router.GET("/albums", getAlbums) -> http://localhost:8080/albums