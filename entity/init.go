package entity

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DB mysql
var DB *gorm.DB

//Cache redis连接
var Cache *redis.Client

//InitDB 初始mysql化数据库
func InitDB() {
	//配置Mysql
	db, err := gorm.Open("mysql", os.Getenv("MysqlConnStr"))
	if err != nil {
		fmt.Println(err)
	}

	//赋值全局变量
	DB = db

	//自动迁移
	AutoMigration()
}

//InitCache 初始化redis缓存
func InitCache() {
	db, err := strconv.Atoi(os.Getenv("RedisDB"))
	if err != nil {
		log.Fatal(err)
	}

	//配置redis
	cache := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RedisConnStr"),
		Password: os.Getenv("RedisPassWord"), // no password set
		DB:       db,                         // use default DB
	})

	//测试
	pong, err := cache.Ping().Result()
	fmt.Println(pong, err)

	//赋值全局变量
	Cache = cache
}

//AutoMigration 自动迁移
func AutoMigration() {
	//禁用表明附属
	DB.SingularTable(true)
	//初始化表
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&Student{})
}
