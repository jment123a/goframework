package common

import (
	"encoding/json"
	"fmt"
	"goframework/entity"
	"strconv"
	"time"

	"github.com/segmentio/ksuid"
)

//GetSession 获取Session
func GetSession(key string) (result []byte, err error) {
	//获取session，[]byte
	result, err = entity.Cache.Get("login_" + key).Bytes()
	if err != nil {
		return nil, err
	}

	//返回
	fmt.Println("GetSession：", key, string(result), err)
	return
}

//SetSession 设置Session
func SetSession(val interface{}) (key string, err error) {
	//创建随机key
	key = ksuid.New().String()

	//对象转为字符串
	json, err := json.Marshal(val)
	if err != nil {
		return
	}

	//设置session，返回
	err = entity.Cache.Set("login_"+key, json, 30*time.Minute).Err()
	fmt.Println("SetSession：", key, err)
	return
}

//SetStudentSession 设置学生Session，防止同一账号在多处登录
func SetStudentSession(val entity.Student) (key string, err error) {
	//创建随机key
	key = ksuid.New().String()

	//检查账号是否已登录
	if res, err := entity.Cache.Get("login_key_" + val.Number).Result(); err == nil {
		//清空已经登录的session
		if _, err1 := ClearSession(res); err1 != nil {
			return "", err1
		}
	}

	//session对象转为字符串
	json, err := json.Marshal(val)
	if err != nil {
		return
	}

	//设置session的key
	err = entity.Cache.Set("login_key_"+val.Number, key, 30*time.Minute).Err()
	if err != nil {
		return
	}

	//设置session，返回
	err = entity.Cache.Set("login_"+key, json, 30*time.Minute).Err()
	fmt.Println("SetStudentSession", key, err)
	return
}

//ClearSession 清空Session
func ClearSession(key string) (result string, err error) {
	//删除session，返回
	res, err := entity.Cache.Del("login_" + key).Result()
	if err != nil {
		return "", err
	}

	fmt.Println("ClearSession：", key, res, err)
	return strconv.FormatInt(res, 10), err
}
