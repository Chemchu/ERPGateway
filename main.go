package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	types "github.com/Chemchu/ERPGateway/types"
	request_redirect "github.com/Chemchu/ERPGateway/utils"
)

func getAPI(c *gin.Context) {
	msg := "Saludos! Bienvenidos al APIGateway"
	successful := true

	APIData := types.APIData{Message: &msg, Successful: &successful}
	out, err := json.Marshal(APIData)
	if err != nil {
		panic(err)
	}
	data := string(out)
	APIRes := types.APIResponse{
		Data: &data,
	}

	c.JSON(http.StatusOK, APIRes)
}

func getStats(c *gin.Context) {
	msg := "Stadistics placeholder"
	successful := true
	APIData := types.APIData{Message: &msg, Successful: &successful}
	out, err := json.Marshal(APIData)
	if err != nil {
		panic(err)
	}
	data := string(out)
	APIRes := types.APIResponse{
		Data: &data,
	}

	c.JSON(http.StatusOK, APIRes)
}

func getGraphQL(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	APIRes := request_redirect.RedirectRequest(body)

	c.JSON(http.StatusOK, APIRes)
}

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Some error occured. Err: %s", errEnv)
	}

	router := gin.Default()
	router.GET("/api", getAPI)
	router.GET("/api/stats/*object", getStats)
	router.POST("/graphql", getGraphQL)

	router.Run("localhost:8080")
}
