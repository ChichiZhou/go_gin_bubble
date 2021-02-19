package models

import (
	"errors"
	"go_gin_project/dao"
)

// 此处规定了前端传值只能是 title
// 并且只能是一个 json
type User struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// 增删改查
// 增
// func CreateAUser(todo *Todo) (err error) {
// 	if err := dao.DB.Create(&todo).Error; err != nil {
// 		return err
// 	}
// 	return
// }

// 根据用户名查询用户密码
func FindAUser(user *User) (err error) {
	var resultUser User
	if err := dao.DB.Where("UserName=?", user.UserName).Find(&resultUser).Error; err != nil {
		return err
	} else {
		if resultUser.Password != user.Password {
			return errors.New("密码错误")
		}
	}
	return
}

// // 查
// func FindTodoById(todo *Todo, id string) (err error) {
// 	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
// 		return err
// 	}
// 	return
// }

// // 改
// func SaveTodo(todo *Todo) (err error) {
// 	if err := dao.DB.Save(&todo).Error; err != nil {
// 		return err
// 	}
// 	return
// }

// // 删
// func DeleteOneTodo(id string) (err error) {
// 	if err := dao.DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
// 		return err
// 	}
// 	return
// }
