package main

import (
	"goframework/api"
	"goframework/entity"
	"goframework/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	//读取配置文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//中间件
	//跨域
	r.Use(middleware.Cors())
	//权限控制
	r.Use(middleware.Auth())

	//连接数据库与缓存连接池
	//redis
	entity.InitCache()
	//mysql
	entity.InitDB()

	//api
	//连通性测试
	r.GET("/ping", api.Ping)
	//总体测试
	r.GET("/test", api.Test)

	//Student
	student := r.Group("/student")
	{
		student.GET("checkLogin", api.CheckLogin)
		student.POST("getCurrentUserInfo", api.GetCurrentUserInfo)
		student.GET("logOut", api.LogOut)
	}

	//开启服务器
	r.Run(":8080")
}
