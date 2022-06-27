package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	request "github.com/Chemchu/ERPGateway/request_redirect"
	"github.com/Chemchu/ERPGateway/types"
)

func getAPI(c *gin.Context) {
	msg := "Saludos! Bienvenidos al APIGateway"
	successful := true

	APIData := types.APIResponse{Message: &msg, Successful: &successful}
	out, err := json.Marshal(APIData)
	if err != nil {
		panic(err)
	}
	data := string(out)
	APIRes := types.APIData{
		Data: &data,
	}

	c.JSON(http.StatusOK, APIRes)
}

func getSummary(c *gin.Context) {
	fecha, err := strconv.ParseInt(c.Param("fecha"), 10, 64)
	if err != nil {
		msg := "Formato de fecha erroneo. Debe ser Unix Epoch en milisegundos."
		successful := false
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Message:    &msg,
			Successful: &successful,
		})
		return
	}
	c.JSON(http.StatusOK, request.RequestGetAnalysis(fecha, os.Getenv("ERPANALYSIS_URL")+"api/analytics/summary"))
}

func postRegistro(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	APIRes := request.RedirectRequest(body, os.Getenv("ERPREGISTRATION_URL")+"api/empleados", "POST")
	c.JSON(http.StatusOK, APIRes)
}

func postGraphQL(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	APIRes := request.RedirectRequest(body, os.Getenv("ERPBACK_URL")+"graphql", "POST")
	c.JSON(http.StatusOK, APIRes)
}

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Some error occured. Err: %s", errEnv)
	}

	router := gin.Default()
	router.GET("/api", getAPI)
	router.GET("/api/analytics/summary/:fecha", getSummary)
	router.POST("/api/registro", postRegistro)
	router.POST("/api/graphql", postGraphQL)

	router.Run("0.0.0.0:8080")
	// router.Run("localhost:8080")

	log.Println("¡API Gateway iniciado!")
}
