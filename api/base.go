package api

import (
	"goframework/entity"
	"time"

	"github.com/gin-gonic/gin"
)

//Ping pingtest
func Ping(c *gin.Context) {
	c.JSON(200, "pong")
}

type test struct {
	RedisPing string `json:"redisPing"`
	RedisSet  string `json:"redisSet"`
	RedisGet  string `json:"redisGet"`
	MysqlTest string `json:"mysqlTest"`
}

//Test alltest
func Test(c *gin.Context) {
	var test test
	//redis
	if v, err := entity.Cache.Ping().Result(); err != nil {
		test.RedisPing = err.Error()
	} else {
		test.RedisPing = v
	}
	if v, err := entity.Cache.Set("k1", "v1", 1*time.Second).Result(); err != nil {
		test.RedisSet = err.Error()
	} else {
		test.RedisSet = v
	}
	if v, err := entity.Cache.Get("k1").Result(); err != nil {
		test.RedisGet = err.Error()
	} else {
		test.RedisGet = v
	}

	//mysql
	if err := entity.DB.DB().Ping(); err != nil {
		test.MysqlTest = err.Error()
	} else {
		test.MysqlTest = "OK"
	}

	c.JSON(200, test)
}
