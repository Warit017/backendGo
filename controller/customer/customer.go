package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer GET Method!",
	})
}
func POST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer POST Method!",
	})
}

func PUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer PUT Method!",
	})
}

func DELETE(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer DELETE Method!",
	})
}
