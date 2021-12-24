package controller

import (
	"math/rand"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	//Get the parameter

	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	//Data verification
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":	422,
			"msg":	"The phone num must be 11 digits!",
		})
	}

	//verification password
	if len(password) < 6 {
		ctx.JSON(http.StatusOK,gin.H{
			"code":	422,
			"msg":	"Password cannot be less than 6 digits!",
		})
	}

	//if not input the name, to a 10 digits random string
	if(len(name) == 0) {
		name = RandomString(10)
	}

	log.Println(name, password, phone)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Register success",		
	})
}

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for k, _:= range result {
		result[k] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}