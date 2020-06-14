package routers

import (
	"github.com/cyberknight01/ck_bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	//告诉gin框架模版文件引用的静态文件去哪里找
	router.Static("/static","static")
	//告诉gin框架去哪里找模版文件
	router.LoadHTMLGlob("templates/*")
	//
	router.GET("/", controller.Indexhandler)

	//v1版本的api
	v1Group := router.Group("/v1")

	//业务逻辑先写出来步骤
	//代办事项逻辑——添加事项/查看事项/修改事项/删除事项
	//1.添加事项，注意每个事项的请求方式
	v1Group.POST("/todo", controller.CreateTodo)
	//2.查看所有的代办事项
	v1Group.GET("/todo", controller.GetTodoList)
	//查看某一个代办事项，七米教程没有实现，课后可以自己试试看
	//v1Group.GET("/todo/:id", func(context *gin.Context) {
	//})
	//3.修改
	v1Group.PUT("/todo/:id", controller.UpdateTodo)
	//4.删除
	v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	return router
}