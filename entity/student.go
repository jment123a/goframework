package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Student 学生类
type Student struct {
	gorm.Model
	Number   string
	Name     string
	Major    string
	Class    string
	PassWord string
	AddTime  *time.Time
	IsExpert bool
}
