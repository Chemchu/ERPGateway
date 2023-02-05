package main

import (
	"fmt"
	"log"
	"os"

	endpoints "github.com/Chemchu/ERPGateway/endpoints_handlers"
	endpoints_summary "github.com/Chemchu/ERPGateway/endpoints_handlers/summary"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Some error occured. Err: %s", errEnv)
	}

	router := gin.Default()
	router.GET("/", endpoints.GetAPI)
	router.GET("/api", endpoints.GetAPI)
	router.GET("/api/analytics/ventas/summary/:fecha", endpoints_summary.GetVentasSummary)
	router.POST("/api/analytics/productos/summary/:fecha", endpoints_summary.GetProductosSummary)
	router.POST("/api/registro", endpoints.PostRegistro)
	router.GET("/api/registro/confirmacion/:token", endpoints.GetRegistroConfirmacion)
	router.POST("/api/registro/confirmacion/:token", endpoints.PostRegistroConfirmacion)
	router.POST("/api/graphql", endpoints.PostGraphQL)
	router.POST("/api/recommender", endpoints.PostRecommendation)

	router.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("GATEWAY_PORT")))

	log.Println("¡API Gateway iniciado!")
}
