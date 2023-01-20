package endpoints

import (
	"io/ioutil"
	"net/http"
	"os"

	request "github.com/Chemchu/ERPGateway/request_redirect"
	"github.com/gin-gonic/gin"
)

func PostGraphQL(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	APIRes := request.RedirectRequest(body, os.Getenv("ERPBACK_URL")+"graphql", "POST")
	c.JSON(http.StatusOK, APIRes)
}
