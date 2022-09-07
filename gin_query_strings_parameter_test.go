package belajar_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/helper"
	"net/http"
	"testing"
)

func TestQueryStringsParameter(t *testing.T) {
	r := gin.Default()

	r.GET("/get/querystring", func(context *gin.Context) {
		firstName := context.Query("first_name")
		lastName := context.Query("last_name")

		context.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})

	err := r.Run()
	helper.PanicIfError(err)

}
