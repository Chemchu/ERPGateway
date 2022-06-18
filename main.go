package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Chemchu/ERPGateway/request_redirect"
	"github.com/Chemchu/ERPGateway/types"
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

func PostRegistro(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	APIRes := request_redirect.RedirectRequest(body, os.Getenv("ERPREGISTRATION_URL")+"api/empleados", "POST")
	c.JSON(http.StatusOK, APIRes)
}

func PostGraphQL(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	APIRes := request_redirect.RedirectRequest(body, os.Getenv("ERPBACK_URL")+"graphql", "POST")
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
	router.POST("/api/registro", PostRegistro)
	router.POST("/api/graphql", PostGraphQL)

	router.Run("localhost:8080")
}
