package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"todo_list/model"
	"todo_list/serializer"
	utils2 "todo_list/utils"
)

type UserService struct {
	UserName string `form:"username" json:"user_name" binding:"required,min=3,max=16"`
	PassWord string `form:"password" json:"pass_word" binding:"required,min=5,max=16"` //限制密码长度
}

// Register 注册
func (service *UserService) Register() *serializer.Response{
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?",service.UserName).
		First(&user).Count(&count) //如果count == 1 则证明数据库中存在这个人
	if count == 1{
		return &serializer.Response{
			Status: 400,
			Msg : "该用户已经存在",
		}
	}
	user.UserName = service.UserName
	//密码加密
	err := user.SetPassword(service.PassWord)
	if err != nil{
		return &serializer.Response{
			Status : 400,
			Msg : fmt.Sprintln(err), //将错误信息传回
		}
	}
	//创建用户
	if err = model.DB.Create(&user).Error ; err != nil{
		return &serializer.Response{
			Status: 500,
			Msg: "数据库创建用户失败",
		}
	}
	return &serializer.Response{
		Status: http.StatusOK,
		Msg: "用户注册成功",
	}
}

// Login 登录
func (service *UserService) Login() *serializer.Response{
	var user model.User
	err := model.DB.Model(&model.User{}).Where("user_name=?",service.UserName).First(&user).Error
	if err != nil{
		if gorm.IsRecordNotFoundError(err){ //没有查询到该用户
			return &serializer.Response{
				Status: http.StatusBadRequest,
				Msg:    "用户不存在，请先注册",
			}
		}
		//如果不是用户不存在而是其他因素
		return &serializer.Response{
			Status: 500,
			Msg: "数据库错误",
		}
	}
	//temp := user.CheckPassword(service.PassWord)
	//if user.UserName == service.UserName && temp { //只有用户名和密码都正确才可以登录
	//	return &serializer.Response{
	//		Status: http.StatusOK,
	//		Msg:    "登录成功",
	//	}
	//}
	//return &serializer.Response{
	//	Status: http.StatusBadRequest,
	//	Msg: "登陆失败,请先注册",
	//}
	if !user.CheckPassword(service.PassWord){
		return &serializer.Response{
			Status: http.StatusBadRequest,
			Msg: "密码错误",
		}
	}
	//发一个token，为了其它功能需要身份验证给前端用来存储的
	//创建一个备忘录，需要token验证，因为需要知道这个是创建的
	token , err := utils2.GenerateToken(user.ID ,service.UserName ,service.PassWord )
	if err != nil{
		return &serializer.Response{
			Status: 500,
			Msg: fmt.Sprint(err) + "token签发错误",
		}
	}
	return &serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User : serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登录成功",
	}
}


