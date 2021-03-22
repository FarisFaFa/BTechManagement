package mysqlfile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type employee_info struct {
	Name      string `xorm:"varchar(20) notnull 'name'" json:"name" form:"Name" binding:"required" `
	Email     string `xorm:"varchar(20) notnull 'email'" json:"email" form:"Email" binding:"required" `
	Position  string `xorm:"varchar(10) notnull 'position'" json:"position" form:"Position" binding:"required" `
	Mobile    string `xorm:"varchar(10) notnull 'mobile'" json:"mobile" form:"Mobile" binding:"required" `
	AccountId int    `xorm:"int unsigned notnull 'account_id'" json:"account_id" form:"Account_id"`
}

type login_info struct {
	Email     string `xorm:"varchar(20) notnull 'email'" json:"email" form:"email" binding:"required" `
	Position  string `xorm:"varchar(10) notnull 'position'" json:"position" form:"position" binding:"required" `
	AccountId int    `xorm:"int unsigned notnull 'account_id'" json:"account_id" form:"account_id"`
	Password  string `xorm:"varchar(32) notnull 'password'" json:"password" form:"password"`
}

// 调取所有用户数据
func XormGetAllEmployee(c *gin.Context) {
	//allTasks := make(map[int64]Task1)
	var allEmployee []employee_info
	err := x.Find(&allEmployee)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, allEmployee)
}

// 插入用户
func XormInsertEmployee(c *gin.Context) {
	// 声明接受的变量
	var NewEmployee employee_info
	// 绑定表单传回的数据Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&NewEmployee); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 首先根据邮箱和职位添加登录信息
	//if err := InsertLoginInfo(NewEmployee); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	insertSQL := "insert into login_info(email,position) values (?,?)"
	_, err := x.Exec(insertSQL, NewEmployee.Email, NewEmployee.Position)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}
	// 根据邮箱查询登录信息，得到登录信息的结构体变量
	NewLogin, err := GetAccountId(NewEmployee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}
	NewEmployee.AccountId = NewLogin[0].AccountId
	// 最后添加员工
	if _, err := x.Insert(NewEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error3": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "添加员工成功")
}

// 插入登录信息
//func InsertLoginInfo(NewEmployee employee_info) (err error) {
//	var NewLogin login_info
//	NewLogin.Email = NewEmployee.Email
//	NewLogin.Position = NewEmployee.Position
//	if _, err = x.Insert(NewLogin); err != nil {
//		return
//	}
//	return
//}

// 查询登录信息
func GetAccountId(NewEmployee employee_info) (NewLogin []login_info, err error) {
	// 根据邮箱，查询登录信息
	err = x.Where("email=?", NewEmployee.Email).Find(&NewLogin)
	if err != nil {
		return
	}
	return
}
