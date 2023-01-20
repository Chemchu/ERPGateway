package endpoints

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func PostRecommendation(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	url := os.Getenv("ERPRECOMMENDER_URL")
	request, reqErr := http.NewRequest("POST", url, bytes.NewBuffer(body))
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
	c.JSON(http.StatusOK, string(resContent))
}
