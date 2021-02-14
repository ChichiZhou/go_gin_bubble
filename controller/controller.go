package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func SayHello(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddTodo(c *gin.Context){
	// 1.从前端页面中把数据拿出来
	var todo Todo
	c.BindJSON(&todo)
	// 2.把数据存入数据库
	// 存入数据的操作是 DB.Create(&todo) 但是这里把存入数据和返回响应写在一起了
	// 3.返回一个响应
	if err := DB.Create(&todo).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func FindAllTodos(c *gin.Context){
	// 查询表中所有的数据
	var todoList []Todo
	if err := DB.Find(&todoList).Error; err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context){
	id := c.Param("id")
	var todo Todo
	if err := DB.Where("id=?", id).First(&todo).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	// 保存更新之后的数据
	c.BindJSON(&todo)
	if err := DB.Save(&todo).Error; err!= nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo (c *gin.Context){
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := DB.Where("id=?", id).Delete(Todo{}).Error;err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{id:"deleted"})
	}
}