package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	dateFormatter "github.com/Chemchu/ERPGateway/date_formatter"
	graphQL "github.com/Chemchu/ERPGateway/graphQL_querys"
	"github.com/Chemchu/ERPGateway/types"
)

func GetSalesAnalysis(fechasParam string, urlEndpoint string) *types.APIResponse {
	msgFallback := "No se ha podido completar el análisis"
	successfulFallback := false
	var failResponse types.APIResponse = types.APIResponse{
		Message:    &msgFallback,
		Successful: &successfulFallback,
	}

	body := types.APIResponse{}

	// Intenta obtener un array de fechas para realizar el analisis.
	fechas := strings.Split(fechasParam, "&")
	switch len(fechas) {
	case 2:
		fechaInicial, err1 := strconv.ParseInt(fechas[0], 10, 64)
		if err1 != nil {
			return &failResponse
		}
		fechaFinal, err2 := strconv.ParseInt(fechas[1], 10, 64)
		if err2 != nil {
			return &failResponse
		}
		body = GetSalesFromDB(fechaInicial, fechaFinal)
	case 1:
		fecha, err1 := strconv.ParseInt(fechas[0], 10, 64)
		if err1 != nil {
			return &failResponse
		}
		body = GetSalesFromDB(dateFormatter.GetStartOfDay(fecha), dateFormatter.GetEndOfDay(fecha))
	default:
		return &failResponse
	}

	var metodoHttp types.HttpMethod = types.POST
	var resContent = FetchWrapper(metodoHttp, urlEndpoint, *body.Data)

	var result types.APIResponse
	json.Unmarshal(resContent, &result)

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

func GetProductsAnalysis(fechasParam string, urlEndpoint string, productsIds []string) *types.APIResponse {
	msgFallback := "No se ha podido completar el análisis"
	successfulFallback := false
	var failResponse types.APIResponse = types.APIResponse{
		Message:    &msgFallback,
		Successful: &successfulFallback,
	}

	body := types.APIResponse{}

	// Intenta obtener un array de fechas para realizar el analisis.
	fechas := strings.Split(fechasParam, "&")
	switch len(fechas) {
	case 2:
		fechaInicial, err1 := strconv.ParseInt(fechas[0], 10, 64)
		if err1 != nil {
			return &failResponse
		}
		fechaFinal, err2 := strconv.ParseInt(fechas[1], 10, 64)
		if err2 != nil {
			return &failResponse
		}
		body = GetSalesFromDB(fechaInicial, fechaFinal)
	case 1:
		fecha, err1 := strconv.ParseInt(fechas[0], 10, 64)
		if err1 != nil {
			return &failResponse
		}
		body = GetSalesFromDB(dateFormatter.GetStartOfDay(fecha), dateFormatter.GetEndOfDay(fecha))
	default:
		return &failResponse
	}

	// TODO: Buscar los productos que le interesan al cliente (realizar un filtrado)
	// for _, venta := range body.Data {
	//
	// }

	var metodoHttp types.HttpMethod = types.POST
	var resContent = FetchWrapper(metodoHttp, urlEndpoint, *body.Data)

	var result types.APIResponse
	json.Unmarshal(resContent, &result)

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
	requestBody, err := json.Marshal(types.GraphQLQuery{
		Query: graphQL.QUERY_VENTAS(),
		Variables: fmt.Sprintf(`
			{
				"find": {
					"fechaInicial": "%d",
					"fechaFinal": "%d"
				},
				"limit": 20000
			}
		`, fechaInicial, fechaFinal),
	})

	if err != nil {
		panic("Error al crear el body de la consulta a DB - GetSalesFromDB")
	}

	var metodoHttp types.HttpMethod = types.POST
	response := FetchWrapper(metodoHttp, os.Getenv("ERPBACK_URL")+"graphql", string(requestBody))

	var result map[string]interface{}
	json.Unmarshal([]byte(response), &result)
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

func FetchWrapper(httpMethod types.HttpMethod, endpointUrl string, body string) []byte {
	// request, reqErr := http.NewRequest(httpMethod.String(), os.Getenv("ERPBACK_URL")+"graphql", bytes.NewBufferString(body))
	request, reqErr := http.NewRequest(httpMethod.String(), endpointUrl, bytes.NewBufferString(body))
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
		fmt.Printf("Client: could not read response body: %s\n", err)
		panic(err)
	}

	defer response.Body.Close()
	return resContent
}
