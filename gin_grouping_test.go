package belajar_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/domain/domain"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/helper"
	"net/http"
	"testing"
)

func TestGrouping(t *testing.T) {
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.GET("/kota/bekasi", func(context *gin.Context) {
			context.JSON(http.StatusOK, domain.FullNameRequest{
				ID:        1,
				FirstName: "Dimas",
				LastName:  "Febriyanto",
				Address: domain.Address{
					Street:   "Jalan-jalan",
					City:     "Bekasi",
					Province: "Jawa Barat",
					Country:  "Indonesia",
				},
			})
		})
		v1.GET("/kota/jakarta", func(context *gin.Context) {
			context.JSON(http.StatusOK, domain.FullNameRequest{
				ID:        1,
				FirstName: "Dimas",
				LastName:  "Febriyanto",
				Address: domain.Address{
					Street:   "Jalan-jalan",
					City:     "Bekasi",
					Province: "Jawa Barat",
					Country:  "Indonesia",
				},
			})
		})
	}

	err := router.Run()
	helper.PanicIfError(err)

}
