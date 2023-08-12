package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		c.Set("request", "中间件")
		// 执行中间件
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("中间件执行完毕", t2)
	}
}
func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare())
	{
		r.GET("middleware", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(http.StatusOK, gin.H{
				"message": req,
			})
		})
		// 局部中间件
		r.GET("middleware2", MiddleWare(), func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(http.StatusOK, gin.H{
				"message": req,
			})
		})
	}
	r.Run(":8000")
}
