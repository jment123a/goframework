package api

import (
	"goframework/model"
	"goframework/service"

	"github.com/gin-gonic/gin"
)

//CheckLogin 检查登录
func CheckLogin(c *gin.Context) {
	//初始化
	var checkLoginModel service.CheckLoginModel

	//验证参数
	if err := c.ShouldBind(&checkLoginModel); err != nil {
		c.JSON(200, model.Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	//检查登录
	key, err := checkLoginModel.CheckLogin()
	if err != nil {
		c.JSON(200, &model.Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	//存储cookie
	c.SetCookie("loginInfo", key, 1, "/", "localhost", false, true)

	c.JSON(200, &model.Response{
		Success: true,
		Data:    key,
	})
	return
}

//GetCurrentUserInfo 获取用户信息
func GetCurrentUserInfo(c *gin.Context) {
	//获取cookie
	key, err := c.Cookie("loginInfo")
	if err != nil {
		c.JSON(200, &model.Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	//获取student
	student, err := service.GetCurrentUserInfo(key)
	if err != nil {
		c.JSON(200, &model.Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	//返回
	c.JSON(200, &model.Response{
		Success: true,
		Data:    student,
	})
	return
}

//LogOut 登出
func LogOut(c *gin.Context) {
	//获取cookie
	key, _ := c.Cookie("loginInfo")

	//删除cookie
	c.SetCookie("loginInfo", "", -1, "/", "localhost", false, true)

	//退出
	service.LogOut(key)

	//返回
	c.JSON(200, &model.Response{
		Success: true,
	})
}
