package belajar_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestTemplates(t *testing.T) {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index", gin.H{
			"title": "Dimas Febriyanto",
			"data": []string{
				"Dimas Febriyanto", "Amanda", "Ucok", "Baba",
			},
		})
	})

	r.Run()
}
