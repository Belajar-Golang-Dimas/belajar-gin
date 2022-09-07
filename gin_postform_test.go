package belajar_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/domain/domain"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/domain/web"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/helper"
	"net/http"
	"testing"
)

func TestPostForm(t *testing.T) {
	r := gin.Default()

	r.POST("/post/postform", func(context *gin.Context) {
		firstName := context.PostForm("first_name")
		lastName := context.PostForm("last_name")

		context.JSON(http.StatusOK, web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data: domain.FullNameRequest{
				ID:        1,
				FirstName: firstName,
				LastName:  lastName,
			},
		})
	})

	err := r.Run()
	helper.PanicIfError(err)
}
