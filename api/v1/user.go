package v1

import (
	"fmt"
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var code int

//查询用户是否存在
func UserExist(c *gin.Context) {}

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	fmt.Println("data.Username", data.Username)
	fmt.Println("code", code)
	if code == errmsg.SUCCSE {
		fmt.Println("创建用户1", code)
		fmt.Println("创建用户2", &data)
		model.CreateUser(&data)
	}

	if code == errmsg.ERROR_USERNAME_USED {
		fmt.Println("用户已存在", code)
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户
func GetUser(c *gin.Context) {}

// 查询用户列表
func GetUsers(c *gin.Context) {}

// 编辑用户
func EditUser(c *gin.Context) {}

// 删除用户
func DeleteUser(c *gin.Context) {}
