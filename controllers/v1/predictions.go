package v1

import (
	"github.com/3dxp/simple-example-api/internal/weather"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPredictions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, weather.GetPredictions())
}

func GetPredictionByLocationID(c *gin.Context) {
	locationID := c.Param("location_id")
	c.IndentedJSON(http.StatusOK, weather.GetPredictionByLocationID(locationID))
}
