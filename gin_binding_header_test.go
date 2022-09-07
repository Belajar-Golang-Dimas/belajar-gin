package belajar_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type HeaderAuth struct {
	XApiKey string `header:"X-API-KEY" binding:"required"`
}

func TestBindingHeader(t *testing.T) {
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var header HeaderAuth
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if header.XApiKey != "RAHASIA" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
