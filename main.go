package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	backendUrl := os.Getenv("ERPBACK_URL")

	request, reqErr := http.NewRequest("POST", backendUrl+"graphql", bytes.NewBufferString(string(body)))
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
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Some error occured. Err: %s", errEnv)
	}

	router := gin.Default()
	router.GET("/api", getAPI)
	router.POST("/graphql", getGraphQL)

	router.Run("localhost:8080")
}
