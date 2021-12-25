package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zxers/zx-gin-vue/common"
	"github.com/zxers/zx-gin-vue/dao"
	"github.com/zxers/zx-gin-vue/db"
	"github.com/zxers/zx-gin-vue/model"
	"github.com/zxers/zx-gin-vue/util"
	"golang.org/x/crypto/bcrypt"
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
		name = util.RandomString(10)
	}

	log.Println(name, password, phone)

	if dao.IsPhoneExist(phone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":	"User exist!",
		})
		return
	}
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg": "加密错误",
		})
	}
	user := model.User{
		Name: 		name,
		Password: 	string(hashedPassword),
		Phone: 		phone,
	}
	db.DB.Create(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "Register success",		
	})
}

func Login(ctx *gin.Context) {
	//获取参数
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	//数据验证
	if phone == ""{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg": "Phone num not null",
		})
		return
	}
	
	if password == ""{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg": "Password not null",
		})
		return
	}

	if len(phone) != 11{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg": "Phone num must be 11 digits1",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Password len not less 6 bigits!"})
		return
	}
	//判断手机号是否存在
	var user model.User
	db.DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg": "User not exist!",
		})
		return
	}
	//判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg": "Password err!",
		})
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg": "System err!",
		})
		log.Printf("token generate error: %v", err)
		return
	}  
	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg": "Login success",
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": user,
		},
	})
}