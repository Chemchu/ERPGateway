package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/Chemchu/ERPGateway/types"
	"github.com/gin-gonic/gin"
)

func GetAPI(c *gin.Context) {
	msg := "Saludos! Bienvenidos al APIGateway"
	successful := true

	APIData := types.APIResponse{Message: &msg, Successful: &successful, Data: nil}
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
