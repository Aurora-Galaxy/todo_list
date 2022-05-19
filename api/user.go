package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/service"
)

// UserRegister 用户注册
// @Tags USER
// @Summary 用户注册
// @Produce json
// @Accept json
// @Param data body service.UserService true "register"
// @Success 200 {object} serializer.Response "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500  {object} serializer.Response "{"status":500,"data":{},"msg":{},"error":"error"}"
// @Router /user/register [post]
func UserRegister(c *gin.Context){
	var userRegister service.UserService
	err := c.ShouldBind(&userRegister)
	if err == nil {
		res := userRegister.Register()
		c.JSON(http.StatusOK,res)
	}else{
		c.JSON(http.StatusBadRequest , err)
	}
}

// UserLogin 用户登录
// @Tags USER
// @Summary 用户登录
// @Produce json
// @Accept json
// @Param     data    body     service.UserService true "login"
// @Success 200 {object} serializer.Response "{"status:200,"data":{},"msg":"登陆成功"}"
// @Failure 500 {object} serializer.Response "{"status":500,"data":{},"msg":{},"error":"error"}"
// @Router /user/login [post]
func UserLogin(c *gin.Context){
	var userRegister service.UserService
	err := c.ShouldBind(&userRegister)
	if err == nil {
		res := userRegister.Login()
		c.JSON(http.StatusOK,res)
	}else{
		c.JSON(http.StatusBadRequest , err)
	}
}
