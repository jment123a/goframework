package middleware

import (
	"goframework/model"
	"goframework/service"

	"github.com/gin-gonic/gin"
)

// Auth 权限配置
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//某些页面不需要登录
		url := c.FullPath()
		if url == "/student/checkLogin" {
			c.Next()
			return
		}

		//获取cookie
		key, err := c.Cookie("loginInfo")
		if err != nil {
			c.JSON(200, &model.Response{
				Success: false,
				Msg:     "请重新登录",
			})
			c.Abort()
			return
		}

		//根据cookie获取session
		student, err := service.GetCurrentUserInfo(key)
		if err != nil {
			c.JSON(200, &model.Response{
				Success: false,
				Msg:     "请重新登录",
			})
			c.Abort()
			return
		}

		//判断是否为登录状态
		if student != nil {
			c.Next()
			return
		}

		c.JSON(200, &model.Response{
			Success: false,
			Msg:     "请重新登录",
		})
		c.Abort()
		return

	}
}
