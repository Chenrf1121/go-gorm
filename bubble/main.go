package main

import (
	"bubble/dao"
	"bubble/model"
	"bubble/router"
)

func main() {
	//创建数据库
	//连接数据库
	err := dao.InitMySql()
	if err!=nil{
		panic(err)
	}
	defer dao.Close()
	//绑定模型
	dao.DB.AutoMigrate(&model.Todo{})

	r := router.SetupRouter()
	r.Run(":8080")
}
