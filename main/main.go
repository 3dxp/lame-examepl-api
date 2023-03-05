package main

import (
	"github.com/3dxp/simple-example-api/weather"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPredictions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, weather.GetPredictions())
}

func getPredictionByLocationID(c *gin.Context) {
	locationID := c.Param("location_id")
	c.IndentedJSON(http.StatusOK, weather.GetPredictionByLocationID(locationID))
}

func main() {
	router := gin.Default()
	apiV1 := router.Group("api/v1")
	{
		predictions := apiV1.Group("/predictions")
		{
			predictions.GET("/", getPredictions)
			predictions.GET("/:location_id", getPredictionByLocationID)
		}
	}

	err := router.Run(":8080")
	if err != nil {
		println(err)
	}
}
