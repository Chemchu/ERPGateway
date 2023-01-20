package endpoints

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	request "github.com/Chemchu/ERPGateway/request_redirect"
)

func PostRegistro(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	APIRes := request.RedirectRequest(body, os.Getenv("ERPREGISTRATION_URL")+"api/empleados", "POST")
	c.JSON(http.StatusOK, APIRes)
}

func GetRegistroConfirmacion(c *gin.Context) {
	token := c.Param("token")
	url := os.Getenv("ERPREGISTRATION_URL") + "api/confirmacion/" + token
	request, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		fmt.Printf("The HTTP request failed with error %s\n", reqErr)
		panic(reqErr)

	}
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		panic(err)
	}

	defer response.Body.Close()

	resContent, _ := ioutil.ReadAll(response.Body)
	c.Data(http.StatusOK, "text/html; charset=utf-8", resContent)
}

func PostRegistroConfirmacion(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	url := os.Getenv("ERPREGISTRATION_URL") + "api/confirmacion/" + c.Param("token")
	jsonData := fmt.Sprintf(`{"password": "%s"}`, strings.Split(string(body), "=")[1])
	request, reqErr := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if reqErr != nil {
		fmt.Printf("The HTTP request failed with error %s\n", reqErr)
		panic(reqErr)
	}
	request.Header.Set("Content-type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		panic(err)
	}

	defer response.Body.Close()

	resContent, _ := ioutil.ReadAll(response.Body)
	c.Data(http.StatusOK, "text/html; charset=utf-8", resContent)
}
