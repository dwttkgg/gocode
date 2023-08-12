package main

import (
	"fmt"
	"ginbook/model"
	"ginbook/uitl"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := uitl.InitDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	// 加载所有页面
	r.LoadHTMLGlob("templates/*")
	r.GET("book/list", booklistHandler)
	r.GET("book/new", booknewHandler)
	r.Run(":8000")
}

func booklistHandler(c *gin.Context) {
	booklist, err := uitl.List()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": err,
		})
		return
	}
	// 返回数据
	c.HTML(http.StatusOK, "book.html", gin.H{
		"code": 0,
		"data": booklist,
	})
	return
}
func booknewHandler(c *gin.Context) {
	var newbook model.Book
	err := c.Bind(&newbook)
	if err != nil {
		fmt.Println("bind book model failed:", err)
	}
	err = uitl.InsertBook(newbook.Title, newbook.Price)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": "插入成功",
	})
}
