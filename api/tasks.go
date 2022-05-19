package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/service"
	utils2 "todo_list/utils"
)

// CreateTask 创建备忘录
// @Tags TASK
// @Summary 创建任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.CreateTaskService true "title"
// @Success 200 {object} serializer.Response "{"status:200,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"msg":{},"error":"error"}
// @Router /task [post]
func CreateTask(c *gin.Context){
	var createTask service.CreateTaskService
	claim , _ := utils2.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask) ; err == nil{  //shouldBind将服务端返回的数据绑定到传入的参数中
		res := createTask.Create(claim.Id)
		c.JSON(http.StatusOK,res)
	}else{
		//logging.Error(err)
		c.JSON(400,err)
	}
}

// ShowTask 展示备忘录
// @Tags TASK
// @Summary 展示任务详细信息
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.ShowTaskService true "show"
// @Success 200 {object} serializer.Response "{"status:200,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"msg":{},"error":"error"}
// @Router /task/:id [get]
func ShowTask(c *gin.Context){
	var showTask service.ShowTaskService
	//claim , _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showTask) ; err == nil{
		res := showTask.Show(c.Param("id")) //c.Param("id")从URL中获取
		c.JSON(http.StatusOK,res)
	}else{
		//logging.Error(err)
		c.JSON(400,err)
	}
}

// ListTasks 展示所有备忘录
// @Tags TASK
// @Summary 获取任务列表
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.ListTasksService true "list"
// @Success 200 {object} serializer.Response "{"status:200,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"msg":{},"error":"error"}
// @Router /tasks [get]
func ListTasks(c *gin.Context){
	var listTasks service.ListTasksService
	claim , _ := utils2.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTasks) ; err == nil{
		res := listTasks.List(claim.Id)
		c.JSON(http.StatusOK,res)
	}else{
		//logging.Error(err)
		c.JSON(400,err)
	}
}

// UpdateTask 更新备忘录
// @Tags TASK
// @Summary 修改任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param	data	body service.UpdateTaskService true "update"
// @Success 200 {object} serializer.Response "{"status:200,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"msg":{},"error":"error"}
// @Router /update/:id [put]
func UpdateTask(c *gin.Context){
	var updateTask service.UpdateTaskService
	//claim , _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateTask) ; err == nil{
		res := updateTask.Update(c.Param("id"))
		c.JSON(http.StatusOK,res)
	}else{
		c.JSON(400,err)
	}
}

// SearchTask 模糊查询，只能查询关键字
// @Tags TASK
// @Summary 查询任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.SearchTaskService true "search"
// @Success 200 {object} serializer.Response "{"status:200,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"msg":{},"error":"error"}
// @Router /search [post]
func SearchTask(c *gin.Context){
	var searchTask service.SearchTaskService
	claim , _ := utils2.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask) ; err == nil{
		res := searchTask.Search(claim.Id)
		c.JSON(http.StatusOK,res)
	}else{
		//logging.Error(err)
		c.JSON(400,err)
	}
}

// DeleteTask 删除某条备忘录
// @Tags TASK
// @Summary 删除任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.DeleteTaskService true "delete"
// @Success 200 {object} serializer.Response "{"status:200,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"msg":{},"error":"error"}
// @Router /delete/:id [delete]
func DeleteTask(c *gin.Context){
	var deleteTask service.DeleteTaskService
	//claim , _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteTask) ; err == nil{
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(http.StatusOK,res)
	}else{
		c.JSON(400,err)
	}
}