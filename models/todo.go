package models

import (
	"github.com/cyberknight01/ck_bubble/dao"
)

//放置模型（即数据库的表定义）和数据库相关的增删改查操作


//建立一个表 to do model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

//对todos 表的增删改查操作
func CreateATodo(todo *Todo) (err error){
	err = dao.DB.Create(&todo).Error
	return err
}

func GetAllTodo()(todolist []*Todo,err error) {
	if err = dao.DB.Find(&todolist).Error; err != nil{
		return nil,err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?",id).First(todo).Error;err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error){
	err = dao.DB.Save(&todo).Error
	return err
}

func DeleteATodo(id string)  (err error){
	err = dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return err
}
