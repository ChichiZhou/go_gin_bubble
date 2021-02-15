package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func ConnectDB()(err error){
	DB, err = gorm.Open("mysql", "root:zhouchichi@(localhost)/bubble?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		return
	}
	err = DB.DB().Ping()
	return err
}
