package models

import "go_gin_project/dao"

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

// 增删改查
// 增
func CreateATodo(todo *Todo) (err error){
	if err := dao.DB.Create(&todo).Error; err != nil{
		return err
	}
	return
}

// 查
func FindAllTodos(todoList *[]Todo) (err error){
	if err := dao.DB.Find(&todoList).Error; err!= nil{
		return err
	}
	return
}

// 查
func FindTodoById(todo *Todo, id string) (err error){
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil{
		return err
	}
	return
}

// 改
func SaveTodo(todo *Todo) (err error){
	if err := dao.DB.Save(&todo).Error; err!= nil{
		return err
	}
	return
}

// 删
func DeleteOneTodo(id string) (err error){
	if err := dao.DB.Where("id=?", id).Delete(Todo{}).Error;err!=nil{
		return err
	}
	return
}

