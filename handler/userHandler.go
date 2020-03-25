package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由方式：/user/{name}
func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

// 路由方式：/user?name=eric&age=18
func UserSaveByQuery(context *gin.Context) {
	username := context.Param("name")
	age := context.Query("age")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}

func UserRegister(context *gin.Context) {
	email := context.PostForm("email")
	password := context.DefaultPostForm("password", "Wa123456")
	passwordAgain := context.DefaultPostForm("password-again", "Wa123456")
	println("email", email, "password", password, "password again", passwordAgain)
}
