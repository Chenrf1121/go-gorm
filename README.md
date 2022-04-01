# go-gorm
goweb开发

## 创建数据库，连接数据库
dsn = 用户名：密码@tcp模式(ip地址加端口号)/数据库名？charset=utf8mb4&parseTime=True&loc=Local
gorm.Open(mysql,dsn)

## 绑定模型
DB.aUTOmIGRATE(模型结构体)

## 使用gin设置路由
1. 告诉框架模板引用的静态文件去哪找 r.Static()
2. 告诉gin去哪找模板文件 r.LoadHTMLGlob()
3. 读取html文件
4. 设置路由组
5. 添加待办事项（增删改查）
### 前端页面填写代办事项，点击提交，发出请求到这
1. 从请求中把数据读出来 c.BindJSON()
2. 存入数据库 DB.Create()
### 查看所有数据
DB.Find()
### 更新数据
1. 通过网页传回的id获得需要更新的id c.Params.Get
2. 通过ID找到数据库中的数据，然后进行修改DB.where("id=?",id).first(&todo),DB.save(&todo)
### 删除数据
通过id删除数据DB.where("id=?",id).Delete()
