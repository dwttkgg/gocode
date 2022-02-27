package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang",
	})
}
func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 使用户访问/hello路径时 执行 sayHello函数
	r.GET("/hello", sayHello)
	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})
	// 启动服务
	r.Run(":9090")
}
