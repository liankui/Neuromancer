package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Neuromancer/model"
	"log"
	"net/http"
)

// 路由方式：/user/{name}
func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

// 路由方式：/user?name=eric&age=18
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.Query("age")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Panicln("err ->", err.Error())
		context.String(http.StatusBadRequest, "输入的数据不合法")
	}
	id := user.Save()
	log.Println("id is ", id)
	context.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(context *gin.Context) {
	var user model.UserModel
	if err := context.Bind(&user); err != nil {
		log.Panicln("login 绑定错误:", err.Error())
	}
	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": u.Email,
		})
	}
}
