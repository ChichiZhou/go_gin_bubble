package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_gin_project/controller"
)

//type Todo struct {
//	ID int `json:"id"`
//	Title string `json:"title"`
//	Status bool `json:"status"`
//}

//func sayHello(c *gin.Context){
//	c.HTML(http.StatusOK, "index.html", nil)
//}
//
//func addTodo(c *gin.Context){
//	// 1.从前端页面中把数据拿出来
//	var todo Todo
//	c.BindJSON(&todo)
//	// 2.把数据存入数据库
//	// 存入数据的操作是 DB.Create(&todo) 但是这里把存入数据和返回响应写在一起了
//	// 3.返回一个响应
//	if err := DB.Create(&todo).Error; err != nil{
//		c.JSON(http.StatusOK, gin.H{
//			"error": err,
//		})
//	} else {
//		c.JSON(http.StatusOK, todo)
//	}
//}

//
//func findAllTodos(c *gin.Context){
//	// 查询表中所有的数据
//	var todoList []Todo
//	if err := DB.Find(&todoList).Error; err!= nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//	}else {
//		c.JSON(http.StatusOK, todoList)
//	}
//}

//
//func findTodo(c *gin.Context){
//
//}
//
//func updateTodo(c *gin.Context){
//	id := c.Param("id")
//	var todo Todo
//	if err := DB.Where("id=?", id).First(&todo).Error; err != nil{
//		c.JSON(http.StatusOK, gin.H{
//			"error": err,
//		})
//		return
//	}
//	// 保存更新之后的数据
//	c.BindJSON(&todo)
//	if err := DB.Save(&todo).Error; err!= nil{
//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//	}else{
//		c.JSON(http.StatusOK, todo)
//	}
//}
//
//func deleteTodo (c *gin.Context){
//	id, ok := c.Params.Get("id")
//	if !ok {
//		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
//		return
//	}
//	if err := DB.Where("id=?", id).Delete(Todo{}).Error;err!=nil{
//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//	}else{
//		c.JSON(http.StatusOK, gin.H{id:"deleted"})
//	}
//}

var (
	DB *gorm.DB
)

func connectDB()(err error){
	DB, err = gorm.Open("mysql", "root:zhouchichi@(localhost)/bubble?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		return
	}
	err = DB.DB().Ping()
	return err
}

func main() {
	fmt.Println("------------------THIS IS THE BACKEND------------------")
	// 创建数据库
	// 连接数据库
	err := connectDB()
	if err != nil {
		panic(err)
	}
	defer DB.Close()  // 程序退出关闭数据库连接
	// 模型绑定
	DB.AutoMigrate(&controller.Todo{})

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
	r.Run(":9000")
}
