package v1

import "github.com/gin-gonic/gin"

func GetLiveZ(c *gin.Context) {
	c.Status(204)
}

func GetReadyZ(c *gin.Context) {
	c.Status(204)
}
