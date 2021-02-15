package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin_project/controller"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
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
	return r
}