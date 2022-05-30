package graphqlRedirect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func RedirectQuery(reqData string) {
	// jsonData := map[string]string{
	//     "query": `
	//         {
	//             people {
	//                 firstname,
	//                 lastname,
	//                 website
	//             }
	//         }
	//     `,
	// }
	jsonValue, _ := json.Marshal(reqData)
	request, err := http.NewRequest("POST", "http://localhost:8080/graphql", bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}
