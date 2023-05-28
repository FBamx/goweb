package controllers

import (
	"net/http"

	"fmt"
	"github.com/gin-gonic/gin"
)

type user struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func Login(ctx *gin.Context) {
	var u user
	if err := ctx.BindJSON(&u); err != nil {
		panic("panic")
	}
	email := u.Email
	password := u.Password
	fmt.Println(email)
	fmt.Println(password)
	ctx.JSON(http.StatusOK, gin.H{
		"email":    email,
		"password": password,
		"message":  "nihao",
	})
}

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "test!!",
	})
}
