package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// // album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// // albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// // getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func getAPI(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, {})
// }

func getGraphQL(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	request, err := http.NewRequest("POST", "http://localhost:8080/graphql", bytes.NewBuffer(body))
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	fmt.Println(*response)

	dbRespone, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(dbRespone))

	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	//router.GET("/api", getAPI)
	router.POST("/graphql", getGraphQL)

	router.Run("localhost:9090")
}
