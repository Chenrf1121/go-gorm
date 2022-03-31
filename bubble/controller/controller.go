package controller

import (

	"bubble/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	url    -->controller  -->logic --> model
	请求来了    控制器       业务逻辑    模型层的怎删改查
*/


func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",nil)
}
func CreateATodo(c *gin.Context) {
	//前端页面填写代办事项，点击提交，发出请求到这
	//1.从请求中把数据拿出来
	var todo model.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	err := model.CreateAtodo(&todo)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,todo)
	}
	//3.返回响应
}
func GetTodoList(c *gin.Context) {
	//查看所有数据
	var todoList []*model.Todo
	todoList,err := model.GetTodoList()
	if err !=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,todoList)
	}
}
func UpdataTodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{"error":"无效id"})
		return
	}
	var todo model.Todo
	err := model.Updata(id,&todo)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}
	c.BindJSON(&todo)
	err = model.Save(&todo)
	if err !=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,todo)
	}
}
func DeleteTodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{
			"error":"无效ID",
		})
		return
	}
	err := model.Delete(id)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,gin.H{id:"deleted"})
	}
}