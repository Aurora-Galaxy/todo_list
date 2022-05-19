package model

//将定义好的结构体各个字段映射到数据库中进行迁移
func migration(){
    //自动迁移模式
	DB.Set("gorm:table_options","charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})
	DB.Model(&Task{}).AddForeignKey("uid","User(id)","CASCADE","CASCADE") //CASCADE跟随外键改动

}