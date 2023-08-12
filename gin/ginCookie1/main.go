package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "notSet"
			// name, value string
			// maxAge int 过期时间,单位为s
			// path cookie所在目录
			// domain
			// secure
			// httpOnly bool
			c.SetCookie("key_cookie", "cookie_value", 60, "/", "localhost", false, true)
		}
		fmt.Println("cookie 的值是:", cookie)
	})
	r.Run(":8000")
}
