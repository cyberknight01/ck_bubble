package main

import (
	"github.com/cyberknight01/ck_bubble/dao"
	"github.com/cyberknight01/ck_bubble/models"
	"github.com/cyberknight01/ck_bubble/routers"
)

func main() {
	//创建数据库，可在sql客户端创建
	//连接数据库，一般程序启动先连接数据库
	if err := dao.InitMysql();err != nil {
		panic(err)
	}
	defer dao.CloseDB()
	//模型和数据库中的表绑定
	dao.DB.AutoMigrate(&models.Todo{})  //表todos
	//路由处理
	router := routers.SetUpRouter()
	//运行服务器
	router.Run(":8080")

}
