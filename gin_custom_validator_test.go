package belajar_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/organizations/Belajar-Golang-Dimas/repositories/new/helper"
	"net/http"
	"testing"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func TestCustomValidator(t *testing.T) {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		name := c.PostForm("name")
		username := c.PostForm("username")
		password := c.PostForm("password")

		user := User{
			Id:       1,
			Name:     name,
			Username: username,
			Password: password,
		}

		validate := new(validator.Validate)
		err := validate.Struct(user)
		helper.PanicIfError(err)

		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"status": "OK",
			"data":   user,
		})
	})
	router.Run(":8080")
}
