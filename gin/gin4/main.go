package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	// bindin:"required" 修饰字段,表示必须存在,传递空值则报错
	User     string `form:"username" json:"user" xml:"user" uri:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" uri:"password" binding:"required"`
}

// 请求访问 curl http://localhost:8000/loginJSON -H 'content-type:application/json' -d '{"user":"root","password":"admin"}'
func main() {
	r := gin.Default()
	r.POST("form", func(c *gin.Context) {
		var json Login
		// 第一种 bind默认绑定from格式
		// 根据请求头中的content-type自动解析
		err := c.Bind(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//// 第二种 将request的body中的数据,自动按照json格式解析到结构体
		//err := c.ShouldBindJSON(&json)
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": err.Error(),
		//	})
		//	return
		//}
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 304,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	})
	r.GET("/:user/:password", func(c *gin.Context) {
		var json Login
		// 第二种 将request的body中的数据,自动按照json格式解析到结构体
		err := c.ShouldBindUri(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 304,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	})
	r.Run(":8000")
}
