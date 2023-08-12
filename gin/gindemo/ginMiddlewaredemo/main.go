package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func mytime(c *gin.Context) {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	fmt.Println("执行耗时为:", since)
}
func main() {
	r := gin.Default()
	r.Use(mytime)
	shopping := r.Group("shopping")
	{
		shopping.GET("index", shoppingIndex)
		shopping.GET("home", shoppingHome)
	}
	r.Run(":8000")
}
func shoppingIndex(c *gin.Context) {
	time.Sleep(5 * time.Second)
}
func shoppingHome(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
