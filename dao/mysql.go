package dao

import (
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//初始化一个全局变量
var (
	DB *gorm.DB
)
//连接数据库
func InitMysql()(err error){
	dsn := "root:Zhs123456*#@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB ,err = gorm.Open("mysql",dsn)
	if err != nil {
		log.Println("open database errors")
		panic(err)
	}
	//测试连通性
	return DB.DB().Ping()
}
//关闭数据库函数封装
func CloseDB() {
	DB.Close()
}