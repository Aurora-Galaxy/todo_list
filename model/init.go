package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

func Database(connstring string){
	//fmt.Println(connstring)
	db , err := gorm.Open("mysql" , connstring)
	if err != nil{
		fmt.Println("mysql数据库连接错误")
	}
	fmt.Println("数据库连接成功")
	db.LogMode(true)  //gorm框架自带的日志输出
	if gin.Mode() == "release"{ //如果gin框架是发行版的话就不用打印日志
		db.LogMode(false)
	}
	db.SingularTable(true) //使gorm创建的表名不加 s ，默认情况下gorm会在创建的表后加 s
	db.DB().SetMaxIdleConns(20) //设置连接池
	db.DB().SetMaxOpenConns(100) //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) //最大连接时间
	DB = db //将db复制到全局变量DB中
	migration()
}