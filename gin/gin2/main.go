package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// url获取
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})
	// api获取
	r.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "ganyu")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.POST("/form", func(c *gin.Context) {
		// 接受post表单请求,没有则为默认值
		c.DefaultPostForm("type", "mm")
		// 接受post表单请求
		username := c.PostForm("username")
		password := c.PostForm("password")
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password %s,hobbys:%v", username, password, hobbys))
	})
	// 上传单个文件
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println("from table get failed,", err)
		}
		// 定义上传的位置
		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	// 上传多个文件
	// 限制上传的大小,默认是32M,当前设置为8M
	r.MaxMultipartMemory = 8 << 20
	r.POST("/uploads", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintln("get failed,", err))
		}
		// 获取所有的文件
		files := form.File["files"]
		// 遍历所有文件
		for _, file := range files {
			// 逐个存
			fmt.Println(file.Filename)
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintln("get failed,", err))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf(" uploaded %d files!", len(files)))
	})
	r.Run(":8080")
}
