package controller

import (
	"net/http"

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


	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "Register success",
		"data":gin.H{
			"msg": "Register",
		},
	})
}