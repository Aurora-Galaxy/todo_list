package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	gorm.Model
	UserName string `gorm:"unique"`//unique 唯一，用户名不可以重复
	PasswordDigest string  //密码存储密文
}

// SetPassword 密码的加密
func (user *User) SetPassword(password string) error{
	bytes , err := bcrypt.GenerateFromPassword([]byte(password),12) //cost 加密难度
	if err != nil{
		return err
	}
	user.PasswordDigest = string(bytes) //加密之后的密码
	return nil
}

// CheckPassword 验证密码
func(user *User)CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest),[]byte(password))
	return err == nil
}
