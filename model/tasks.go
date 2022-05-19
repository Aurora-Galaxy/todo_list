package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User User `gorm:"ForeignKey:Uid"`  //设置外键，将用户和任务进行关联
	Uid uint `gorm:"not null"`
	Title string `gorm:"index,not null"` //index 普通索引
	Status uint `gorm:"default:'0'"` // 0 表示未完成
	Content string `gorm:"type:longtext"` //longtext 一种比较长的字符串类型
	StartTime int64 //备忘录的开始时间
	EndTime int64 //备忘录完成时间
}
