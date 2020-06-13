package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

//初始化一个全局变量
var (
	DB *gorm.DB
)

//建立一个表 to do modul
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func InitMysql()(err error){
	dsn := "root:Zhs123456*#@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB ,err = gorm.Open("mysql",dsn)
	if err != nil {
		panic(err)
	}
	//测试连通性
	return DB.DB().Ping()
}

func main() {
	//创建数据库，可在sql客户端创建
	//连接数据库，一般程序启动先连接数据库
	err := InitMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	//模型和数据库中的表绑定
	DB.AutoMigrate(&Todo{})  //表todos

	router := gin.Default()
	//告诉gin框架模版文件引用的静态文件去哪里找
	router.Static("/static","static")
	//告诉gin框架去哪里找模版文件
	router.LoadHTMLGlob("templates/*")
	//
	router.GET("/", func(context *gin.Context) {
		context.HTML(200,"index.html",nil)
	})

	//v1版本的api
	v1Group := router.Group("/v1")


	//业务逻辑先写出来步骤
	//代办事项
	//添加事项，注意每个事项的请求方式
	v1Group.POST("/todo", func(context *gin.Context) {
		//前端页面填写代办事项，点击提交，会发送请求到这里
		//1.从请求中取出数据
		var todo Todo
		context.BindJSON(&todo)
		//2.把数据存入数据库
		if err = DB.Create(&todo).Error;err != nil {
			context.JSON(200,gin.H{"error":err.Error()})
		} else {
			context.JSON(http.StatusOK,todo)
			//context.JSON(http.StatusOK,gin.H{
			//	"code" : 2000,
			//	"msg" : "success",
			//	"data" : todo,
			//})	//此处返回的形式和前端规定有关
		}
		//3.返回响应


	})
	//查看所有的代办事项
	v1Group.GET("/todo", func(context *gin.Context) {
		//查询todos表里的所有数据
		var todolist []Todo
		if err = DB.Find(&todolist).Error;err != nil {
			context.JSON(http.StatusOK,gin.H{
				"find errors" : err.Error(),
			})
		} else {
			context.JSON(http.StatusOK,todolist)

		}

	})
	//查看某一个代办事项，七米教程没有实现，课后可以自己试试看
	//v1Group.GET("/todo/:id", func(context *gin.Context) {
	//
	//})

	//修改
	v1Group.PUT("/todo/:id", func(context *gin.Context) {
		//1.首先从前端获取参数，并根据参数查找数据库
		id,ok := context.Params.Get("id")
		if !ok { context.JSON(http.StatusOK,gin.H{ "get params error" : "id不存在" })
			return
		}
		var todo Todo
		if err = DB.Where("id=?",id).First(&todo).Error;err != nil {
			context.JSON(http.StatusOK,gin.H{"lookup error" : err.Error()})
		}
		//2.然后把前端返回的数据条目给这个todo
		context.BindJSON(&todo)
		//3.在保存在数据库中
		if err = DB.Save(&todo).Error;err != nil {
			context.JSON(http.StatusOK,gin.H{ "save error" : err.Error()})
		} else {
			context.JSON(http.StatusOK,todo)
		}

	})
	//删除
	v1Group.DELETE("/todo/:id", func(context *gin.Context) {
		id, ok := context.Params.Get("id")
		if !ok {
			context.JSON(http.StatusOK,gin.H{ "get params error" : "id不存在" })
			return
		}
		if err = DB.Where("id=?",id).Delete(Todo{}).Error;err != nil {
			context.JSON(http.StatusOK,gin.H{ " deleted error" : err.Error() })
		} else {
			context.JSON(http.StatusOK,gin.H{id:"deleted"})
		}
	})

	//运行服务器
	router.Run(":8080")

}
