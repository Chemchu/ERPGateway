package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	types "erpgateway/types"
	request_redirect "erpgateway/utils"
)

func getAPI(c *gin.Context) {
	msg := "Saludos! Bienvenidos al APIGateway"
	successful := true
	APIRes := types.APIResponse{
		Message:    &msg,
		Successful: &successful,
	}

	c.JSON(http.StatusOK, APIRes)
}

func getGraphQL(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
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
	router.POST("/graphql", getGraphQL)

	router.Run("localhost:8080")
}
