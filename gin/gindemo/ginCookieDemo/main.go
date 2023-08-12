package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// cookie的缺点 不安全,明文,可以被禁用,增加带宽消耗,cookie有上限
// session可以弥补cookie的不足,session必须依赖cookie使用,生成一个sessionid放在cookie中传递给客户端即可
func AuthCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("key_cookie")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "StatusUnauthorized",
			})
			// 中断调用
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
func main() {
	r := gin.Default()
	r.GET("login", func(c *gin.Context) {
		c.SetCookie("key_cookie", "cookie_value", 60, "/", "localhost", false, true)
		c.String(http.StatusOK, "Login sucessful")
	})
	r.GET("home", AuthCookie(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "home",
		})
	})

	r.Run(":8000")
}
