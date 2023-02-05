package endpoints_summary

import (
	"fmt"
	"net/http"
	"os"

	request "github.com/Chemchu/ERPGateway/request_redirect"
	"github.com/gin-gonic/gin"
)

type ProductsIds struct {
	IDS []string `json:"ids"`
}

func GetProductosSummary(c *gin.Context) {
	var productosIds ProductsIds
	if err := c.ShouldBindJSON(&productosIds); err != nil {
		fmt.Printf("Error: %+v\n", err)
		c.JSON(http.StatusOK, gin.H{"message": err, "successful": false})
		return
	}

	c.JSON(http.StatusOK, request.GetProductsAnalysis(c.Param("fecha"), os.Getenv("ERPANALYSIS_URL")+"api/analytics/productos/summary", productosIds.IDS))
}
