package main

import (
	v1 "github.com/3dxp/simple-example-api/controllers/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	apiV1 := router.Group("api/v1")
	{
		prediction := apiV1.Group("/predictions")
		{
			prediction.GET("/", v1.GetPredictions)
			prediction.GET("/:location_id", v1.GetPredictionByLocationID)
		}
		status := apiV1.Group("/status")
		{
			status.GET("livez", v1.GetLiveZ)
			status.GET("readyz", v1.GetReadyZ)
		}
	}

	err := router.Run(":8080")
	if err != nil {
		println(err)
	}
}
