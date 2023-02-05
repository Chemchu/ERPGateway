package endpoints_summary

import (
	"net/http"
	"os"

	request "github.com/Chemchu/ERPGateway/request_redirect"
	"github.com/gin-gonic/gin"
)

func GetVentasSummary(c *gin.Context) {
	c.JSON(http.StatusOK, request.GetSalesAnalysis(c.Param("fecha"), os.Getenv("ERPANALYSIS_URL")+"api/analytics/ventas/summary"))
}
