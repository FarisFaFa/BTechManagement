package main

import (
	"fmt"
	"management/dao/mysqlfile"
	"management/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//var db *sql.DB

func main() {
	// 初始化数据库
	//err := mysqlfile.InitDB()
	err := mysqlfile.XormInit()
	// 判断错误
	if err != nil {
		fmt.Printf("error is %v", err)
	}
	// 控制台打印访问成功
	fmt.Println("success!")
	// 创建路由
	r := gin.Default()
	// 绑定路由规则

	// 登录页面
	r.LoadHTMLGlob("./frontend/*")
	r.Handle("GET", "/", func(context *gin.Context) {
		// 返回HTML文件，响应状态码200，html文件名为index.html，模板参数为nil
		context.HTML(http.StatusOK, "login.html", nil)
	})

	//登陆
	r.POST("/management/loginCheck", mysqlfile.LoginCheck)

	// 导航页
	r.GET("/management/navigation", mysqlfile.PermissionCheck)

	//中间件
	r.Use(middleware.PermissionCheck)
	{
		// 任务管理页面
		r.GET("/management/task_management", mysqlfile.XormGetAllTasks)
		// 插入新任务
		r.POST("/management/task_management/modify", mysqlfile.InsertTask)

		// 用户管理界面
		r.GET("/management/employee_management", mysqlfile.XormGetAllEmployee)
		// 插入新用户
		r.POST("/management/employee_management/new", mysqlfile.XormInsertEmployee)

	}
	// 监听端口
	r.Run(":8080")
}
