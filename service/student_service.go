package service

import (
	"encoding/json"
	"goframework/common"
	"goframework/entity"
	"goframework/model"
)

//CheckLoginModel 登录模型
type CheckLoginModel struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	PassWord string `form:"passWord" json:"passWord" binding:"required"`
}

//CheckLogin 检查登录
func (checkLogin *CheckLoginModel) CheckLogin() (key string, err error) {
	//初始化
	var student entity.Student

	//验证账号密码
	if err = entity.DB.Where("number = ? and pass_Word = ?", checkLogin.UserName, checkLogin.PassWord).First(&student).Error; err != nil {
		return "", err
	}

	//redis存储session
	key, err = common.SetStudentSession(student)
	return
}

//GetCurrentUserInfo 获取用户信息
func GetCurrentUserInfo(key string) (student *entity.Student, err error) {
	//初始化
	var stu entity.Student

	//获取session
	val, err := common.GetSession(key)
	if err != nil {
		return &entity.Student{}, err
	}

	//session转换为student
	err = json.Unmarshal(val, &stu)
	if err != nil {
		return &entity.Student{}, &model.Error{Msg: "结构体转换失败"}
	}
	return &stu, nil
}

//LogOut 退出登录
func LogOut(key string) (err error) {
	//获取student
	student, err := GetCurrentUserInfo(key)
	if err != nil {
		return
	}

	//清空login
	if _, err = common.ClearSession(key); err != nil {
		return
	}

	//清空login_key
	if _, err = common.ClearSession("key_" + student.Number); err != nil {
		return
	}

	return
}
