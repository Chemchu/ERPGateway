package request_redirect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Chemchu/ERPGateway/types"
)

func RedirectRequest(body []byte) *types.APIResponse {
	backendUrl := os.Getenv("ERPBACK_URL")
	request, reqErr := http.NewRequest("POST", backendUrl+"graphql", bytes.NewBufferString(string(body)))
	if reqErr != nil {
		fmt.Printf("The HTTP request failed with error %s\n", reqErr)
		panic(reqErr)
	}
	request.Header.Add("Content-Type", "application/json")
	defer request.Body.Close()

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		panic(err)
	}
	defer response.Body.Close()

	var result map[string]interface{}
	resContent, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(resContent), &result)
	data := result["data"].(map[string]interface{})
	dbData, _ := json.Marshal(data)
	resData := string(dbData)

	APIRes := types.APIResponse{
		Data: &resData,
	}

	return &APIRes
}
