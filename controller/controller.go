package controller

import (
	"fmt"
	"go_gin_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 显示首页
func SayHello(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddTodo(c *gin.Context) {
	// 1.从前端页面中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	fmt.Println(todo)
	// 2.把数据存入数据库
	// 存入数据的操作是 DB.Create(&todo) 但是这里把存入数据和返回响应写在一起了
	// 3.返回一个响应
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func FindAllTodos(c *gin.Context) {
	// 查询表中所有的数据
	var todoList []models.Todo
	if err := models.FindAllTodos(&todoList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := models.FindTodoById(&todo, id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	// 保存更新之后的数据
	c.BindJSON(&todo)
	if err := models.SaveTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteOneTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}

func FindUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	if err := models.FindAUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户不存在或密码错误",
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "用户存在",
		})
	}
}
