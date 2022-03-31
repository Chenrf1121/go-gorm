package model

import "bubble/dao"

//model
type Todo struct {
	ID int`json:"id"`
	Title string`json:"title"`
	Status bool `json:"status"`
}

//todo怎删改查
//创建
func CreateAtodo(todo *Todo)(err error){
	if err = dao.DB.Create(&todo).Error;err!=nil{
		return err
	}
	return
}

func GetTodoList()(todolist []*Todo,err error){

	if err = dao.DB.Find(&todolist).Error;err!=nil{
		return nil,err
	}
	return
}

func Updata(id string,todo *Todo)(err error){
	if err = dao.DB.Where("id=?",id).First(&todo).Error;err!=nil{
		return err
	}
	return
}
func Save(todo *Todo) (err error){
	if err= dao.DB.Save(&todo).Error;err!=nil{
		return err
	}
	return
}
func Delete(id string)(err error){
	if err = dao.DB.Where("id=?",id).Delete(Todo{}).Error;err!=nil{
		return err
	}
	return
}