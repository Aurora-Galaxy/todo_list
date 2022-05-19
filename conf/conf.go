package conf

import (
	"fmt"
	"gopkg.in/ini.v1"  //需要事先安装依赖 go get gopkg.in/ini.v1
	"strings"
	"todo_list/model"
)

var(
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init(){
	file , err := ini.Load("./conf/config.ini")//读取配置文件
	if err != nil{
		fmt.Println("配置文件读取错误，请检查文件路径")
	}
	LoadServer(file)  //加载服务器的配置
	LoadMysql(file)   //加载mysql的配置
	//"user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
    path := strings.Join([]string{DbUser,":",DbPassWord,"@tcp(",DbHost,":",DbPort,")/",DbName,"?charset=utf8mb4&parseTime=True"},"")
    model.Database(path)
}

func LoadServer(file *ini.File){
	AppMode = file.Section("service").Key("AppMode").String() //section选中.ini文件中的service段，Key方法类似
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File){
	Db = file.Section("mysql").Key("Db").String()
	DbHost= file.Section("mysql").Key("DbHost").String()
	DbPort= file.Section("mysql").Key("DbPort").String()
	DbUser= file.Section("mysql").Key("DbUser").String()
	DbPassWord= file.Section("mysql").Key("DbPassWord").String()
	DbName= file.Section("mysql").Key("DbName").String()
}