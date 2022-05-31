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
type APIResponse struct {
	Message    *string `json:"message,omitempty"`
	Data       *string `json:"data,omitempty"`
	Successful *bool   `json:"successful,omitempty"`
}

func getAPI(c *gin.Context) {
	msg := "Saludos! Bienvenidos al APIGateway"
	successful := true
	APIRes := APIResponse{
		Message:    &msg,
		Successful: &successful,
	}
	c.IndentedJSON(http.StatusOK, APIRes)
}

func getGraphQL(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	request, reqErr := http.NewRequest("POST", "http://localhost:9090/graphql", bytes.NewBufferString(string(body)))
	request.Header.Add("Content-Type", "application/json")
	if reqErr != nil {
		fmt.Printf("The HTTP request failed with error %s\n", reqErr)
	}
	defer request.Body.Close()

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()

	resContent, _ := ioutil.ReadAll(response.Body)

	msg := "Éxito al realizar la petición"
	data := string(resContent)
	successful := true
	APIRes := APIResponse{
		Message:    &msg,
		Data:       &data,
		Successful: &successful,
	}

	c.IndentedJSON(http.StatusOK, APIRes)
}

func main() {
	router := gin.Default()
	router.GET("/api", getAPI)
	router.POST("/graphql", getGraphQL)

	router.Run("localhost:8080")
}
