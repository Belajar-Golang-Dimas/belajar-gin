package belajar_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/domain/domain"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/domain/web"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/helper"
	"net/http"
	"testing"
)

func getHandler(context *gin.Context) {
	context.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: domain.FullNameRequest{
			ID:        1,
			FirstName: "Dimas",
			LastName:  "Febriyanto",
		},
	})
}

func postHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "OK",
		"data": map[string]string{
			"first_name": "Dimas",
			"last_name":  "Febriyanto",
		},
	})
}

func TestMains(T *testing.T) {
	r := gin.Default()
	r.GET("/category/get", getHandler)
	r.POST("/category/post", postHandler)

	err := r.Run()
	helper.PanicIfError(err)
}
