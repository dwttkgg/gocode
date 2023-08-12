package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("redirect", func(c *gin.Context) {
		// 重定向
		// 支持内部和外部重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 异步处理
	// 异步时只能使用c 的只读副本,不能使用c本身
	r.GET("/long_async", func(c *gin.Context) {
		copyContext := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行", copyContext.Request.URL.Path)
		}()
	})
	// 同步处理
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行", c.Request.URL.Path)
	})
	r.Run(":8000")
}
