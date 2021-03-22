package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	//"net/http"
)

//var BackToLogin = false

func PermissionCheck(c *gin.Context) {
	c.Next()
	cookie, err := c.Cookie("key_cookie")

	if err != nil {
		fmt.Println("登录信息过期或未登陆")
		//ToLogin(c)
	} else {
		fmt.Println("cookie:", cookie)
		//BackToLogin = false
	}
}

//
//func ToLogin(c *gin.Context) {
//	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/management/login")
//}
