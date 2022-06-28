package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Chemchu/ERPGateway/dateFormatter"
	graphQL "github.com/Chemchu/ERPGateway/graphQLquerys"
	"github.com/Chemchu/ERPGateway/types"
)

func RequestGetAnalysis(fecha int64, service string) *types.APIResponse {
	body := GetSalesFromDB(dateFormatter.GetStartOfDay(fecha), dateFormatter.GetEndOfDay(fecha))
	url := service
	request, reqErr := http.NewRequest("POST", url, bytes.NewBufferString(*body.Data))
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

	resContent, _ := ioutil.ReadAll(response.Body)
	var result types.APIResponse
	json.Unmarshal([]byte(resContent), &result)

	var msg *string = result.Message
	var successful *bool = result.Successful
	var data *string = result.Data

	APIRes := types.APIResponse{
		Message:    msg,
		Successful: successful,
		Data:       data,
	}

	return &APIRes
}

func GetSalesFromDB(fechaInicial int64, fechaFinal int64) types.APIResponse {
	body, err := json.Marshal(types.GraphQLQuery{
		Query: graphQL.QUERY_VENTAS(),
		Variables: fmt.Sprintf(`
			{
				"find": {
					"fechaInicial": "%d",
					"fechaFinal": "%d"
				}
			}
		`, fechaInicial, fechaFinal),
	})

	if err != nil {
		panic("Error al crear el body de la consulta a DB - GetSalesFromDB")
	}

	request, reqErr := http.NewRequest("POST", os.Getenv("ERPBACK_URL")+"graphql", bytes.NewBufferString(string(body)))
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

	resContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}

	defer response.Body.Close()

	var result map[string]interface{}
	json.Unmarshal([]byte(resContent), &result)
	dbData := result["data"].(map[string]interface{})
	dataBytes, err := json.Marshal(dbData)

	if err != nil {
		panic("Error en marshalling")
	}

	var ventasResult map[string][]types.Venta
	json.Unmarshal([]byte(dataBytes), &ventasResult)
	dbVentas := ventasResult["ventas"]
	ventasBytes, err := json.Marshal(dbVentas)

	if err != nil {
		panic("Error en marshalling")
	}

	data := string(ventasBytes)
	APIRes := types.APIResponse{
		Data: &data,
	}

	return APIRes
}
