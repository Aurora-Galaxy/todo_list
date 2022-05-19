package service

import (
	"time"
	"todo_list/model"
	"todo_list/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title"form:"title"`      //标题
	Content string `json:"content" form:"content"` //内容
	Status  int    `json:"status"form:"status"`    //任务的状态
}

type ShowTaskService struct{
}

type ListTasksService struct{
	//分页功能
	//PageNum int `json:"page_num"form:"page_num"`
	//PageSize int `json:"page_size"form:"page_size"`
}

type UpdateTaskService struct{
	Title   string `json:"title"form:"title"`      //标题
	Content string `json:"content" form:"content"` //内容
	Status  uint    `json:"status"form:"status"`    //任务的状态
}

type SearchTaskService struct{
   Info string `json:"info"form:"info"`  //关键字，按照该字查询
}

type DeleteTaskService struct{

}

// Create 增加一条备忘录
func (service *CreateTaskService) Create(userid uint) *serializer.Response {
	var user model.User
	model.DB.First(&user,userid)
	task := model.Task{
		User: user,
		Uid: user.ID,  //用来区分该任务是谁创建的
		Title: service.Title,
		Status: 0, //默认新建的任务为未完成的
		Content: service.Content,
		StartTime: time.Now().Unix(),
		EndTime: 0,
	}
	err := model.DB.Create(&task).Error
	if err != nil{
		return &serializer.Response{
			Status: 500,
			Msg: "创建备忘录失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg: "创建成功",
	}
}

// Show 按照URL中输入的id展示该条备忘录内容
func (service *ShowTaskService) Show(taskId string) *serializer.Response{
	var task model.Task
	err := model.DB.First(&task , taskId).Error
	if err != nil{
		return &serializer.Response{
			Status: 400,
			Msg: "查询失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Data: serializer.BuildTask(task),
	}
}

// List 展示一个用户的所有备忘录内容
func (service *ListTasksService) List(userid uint) *serializer.Response{
	var tasks []model.Task
	count := 0
	err := model.DB.Model(&model.Task{}).Where("uid=?",userid).Find(&tasks).Count(&count).Error //count 计算查询到几条信息
	if err != nil{
		return &serializer.Response{
			Status: 400,
			Msg: "展示失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Data: serializer.BuildTasks(tasks),
		Total: count, //备忘录总数
	}
}

// Update 更新用户的备忘录
func (service *UpdateTaskService) Update(taskid string) *serializer.Response{
	var task model.Task
	err:= model.DB.Model(&model.Task{}).First(&task,taskid).Error
	if err != nil{
		return &serializer.Response{
			Status: 400,
			Msg: "数据查找失败",
		}
	}
	task.Status = service.Status
	task.Content = service.Content
	task.Title = service.Title
	if err = model.DB.Save(&task).Error ; err != nil{
		return &serializer.Response{
			Status: 401,
			Msg: "数据保存失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg: "数据更新成功",
		Data: serializer.BuildTask(task),
	}
}

// Search 查询  需要传入用户id，限制只能查询自己的备忘录内容
func(service *SearchTaskService) Search(userid uint) *serializer.Response{
	var task []model.Task
	count := 0
	err := model.DB.Model(&model.Task{}).Where("uid=?",userid).
		Where("title LIKE ? OR content LIKE ?","%"+service.Info+"%","%"+service.Info+"%"). //模糊查询的sql语句
	Find(&task).Count(&count).Error //count 计算查询到几条信息
	if err != nil{
		return &serializer.Response{
			Status: 400,
			Msg: "查询失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Data: serializer.BuildTasks(task),
		Total: count, //模糊查询符合条件的数量
	}
}

// Delete 删除一条备忘录，只是软删除,在delete_at会显示删除时间，但内容还是会存在
func (service *DeleteTaskService) Delete(taskid string) *serializer.Response{
	var task model.Task
	err := model.DB.Delete(&task,taskid).Error
	if err != nil{
		return &serializer.Response{
			Status: 500,
			Msg: "删除失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg: "删除成功",
	}
}