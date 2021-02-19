package routers

import (
	"go_gin_project/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 解决跨域问题
	r.Use(cors.Default())
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	// 规定了首页是通过 SayHello 渲染的
	r.GET("/", controller.SayHello)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.AddTodo)

		// 查看所有的待办事项
		v1Group.GET("/todo", controller.FindAllTodos)

		//// 查看某一个待办事项
		//v1Group.GET("/todo/:id", func(c *gin.Context) {
		//
		//})

		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	// 由于这里是 v2Group.POST("/", controller.FindUser)
	// 所以在前端需要用 http://localhost/9000/login/
	v2Group := r.Group("login")
	{
		v2Group.POST("/", controller.FindUser)
	}
	return r
}
