package mysqlfile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `xorm:"varchar(20) notnull 'email'" json:"email" form:"Email" binding:"required" `
	Password string `xorm:"varchar(32) notnull 'password'" json:"password" form:"Password" binding:"required"`
}

type Permission struct {
	Position       string `xorm:"varchar(10) notnull 'position'" json:"position"`
	Permission_url string `xorm:"varchar(10) notnull 'permission_url'" json:"permission_url"`
	Permission_id  string `xorm:"varchar(10) notnull 'permission_id'" json:"permission_id"`
}

var check []login_info

func LoginCheck(c *gin.Context) {
	// 声明接受登录信息的变量
	var L Login
	// 绑定表单传回数据，用Bind()
	if err := c.Bind(&L); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"获取登录信息错误": err.Error()})
		return
	}
	// 将登录表重的信息赋值给全局变量
	check = getLoginInfo(L)
	// 判断是否成功
	if L.Password == check[0].Password {
		//c.JSON(http.StatusOK, gin.H{"用户登录成功！邮箱": check[0].Email})
		c.Redirect(http.StatusMovedPermanently, "http://18.191.157.203:8080/management/navigation")
		c.SetCookie("key_cookie", check[0].Email+check[0].Position, 30, "/", "localhost", false, true)
	} else {
		c.JSON(http.StatusOK, gin.H{"用户验证失败！邮箱：": check[0].Email})
		return
	}
}

func PermissionCheck(c *gin.Context) {
	// 定义接收职位的变量
	var position Permission
	position.Position = check[0].Position
	// 定义接收权限的变量
	var PermissionList []Permission
	// xorm根据check的职位，查询权限表重的内容
	if err := x.Where("position=?", position.Position).Find(&PermissionList); err != nil {
		c.JSON(http.StatusOK, gin.H{"权限查找失败": err.Error()})
		return
	}
	c.JSON(http.StatusOK, PermissionList)

}

func getLoginInfo(L Login) (check []login_info) {
	//
	if err := x.Where("email=?", L.Email).Find(&check); err != nil {
		return
	}
	return
}
