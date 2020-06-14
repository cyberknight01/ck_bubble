package controller

import (
	"github.com/cyberknight01/ck_bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// url ——> controller ——> logic ——> models
// 请求来了 --> 控制器 --> 业务逻辑 --> 模型层的增删改查
//controller里只负责调用逻辑，不负责处理逻辑，业务复杂情况会有个
//logic层，此处都放在了models里


func Indexhandler(context *gin.Context) {
	context.HTML(200,"index.html",nil)
}

func CreateTodo(context *gin.Context) {
	//前端页面填写代办事项，点击提交，会发送请求到这里
	//1.从请求中取出数据
	var todo models.Todo
	context.BindJSON(&todo)
	//2.把数据存入数据库
	if err := models.CreateATodo(&todo); err != nil {
		context.JSON(200,gin.H{"error":err.Error()})
	} else {
		//3.返回响应
		context.JSON(http.StatusOK,todo)
	}

}

func GetTodoList(context *gin.Context) {
	//查询todos表里的所有数据
	if todolist, err := models.GetAllTodo(); err != nil {
		context.JSON(http.StatusOK,gin.H{"find errors" : err.Error()})
		} else {
		context.JSON(http.StatusOK,todolist)
	}
}

func UpdateTodo(context *gin.Context) {
	//1.首先从前端获取参数，并根据参数查找数据库
	id,ok := context.Params.Get("id")
	if !ok { context.JSON(http.StatusOK,gin.H{ "get params error" : "id不存在" })
		return
	}
	todo, err := models.GetATodo(id)  //下面要用到todo这个变量，所以不能放在if局部里
	if err != nil {
		context.JSON(http.StatusOK,gin.H{ " mysql get id error" : err.Error()})
		return
	}
	//2.然后把前端返回的数据条目给这个todo
	context.BindJSON(&todo)
	//3.在保存在数据库中
	if err = models.UpdateATodo(todo);err != nil {
		context.JSON(http.StatusOK,gin.H{ "save error" : err.Error()})
	} else {
		context.JSON(http.StatusOK,todo)
	}
}

func DeleteATodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK,gin.H{ "get params error" : "id不存在" })
		return
	}
	if err := models.DeleteATodo(id);err != nil {
		context.JSON(http.StatusOK,gin.H{ " deleted error" : err.Error() })
	} else {
		context.JSON(http.StatusOK,gin.H{id:"deleted"})
	}
}