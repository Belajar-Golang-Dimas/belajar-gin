package belajar_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/helper"
	"testing"
)

func TestMiddleware(t *testing.T) {

	r := gin.New()

	gin.SetMode(gin.ReleaseMode)

	r.Use(gin.Logger())

	r.GET("/middleware", func(context *gin.Context) {
		panic("eror nih bosque")
	})

	err := r.Run()
	helper.PanicIfError(err)

}
