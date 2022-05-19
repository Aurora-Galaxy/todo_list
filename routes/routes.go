package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"todo_list/api"
	_ "todo_list/docs"  //一定要导入，否则在网页会报错
	"todo_list/middleware"
)

func NewRouter() *gin.Engine{
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	r.Use(sessions.Sessions("mysession",store))
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register",api.UserRegister)
		v1.POST("user/login",api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())  //执行以下路由会先执行中间件，验证token
		{
			authed.POST("task",api.CreateTask)
			authed.GET("task/:id",api.ShowTask)
			authed.GET("tasks",api.ListTasks)
			authed.PUT("update/:id",api.UpdateTask)
			authed.POST("search",api.SearchTask)
			authed.DELETE("delete/:id",api.DeleteTask)
	    }
	}
	return r
}