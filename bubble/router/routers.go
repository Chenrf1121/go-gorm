package router

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	//告诉gin框架模板为念引用的静态文件去哪找
	r.Static("/static","static")
	//告诉gin框架去哪找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	//v1
	v1Gropu := r.Group("v1")
	{
		//代办事项
		//添加
		v1Gropu.POST("/todo", controller.CreateATodo)
		//查看
		v1Gropu.GET("todo",controller.GetTodoList)
		v1Gropu.GET("/todo/:id", func(c *gin.Context) {
			//查看某一个事项
		})
		//修改
		v1Gropu.PUT("/todo/:id", controller.UpdataTodo)
		//删除
		v1Gropu.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}