package main

import (
	"net/http"

	"github.com/gin-gonic/gin/testdata/protoexample"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 1 json响应
	r.GET("someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "someJSON",
			"status":  200,
		})
	})
	// 2 struct 响应
	r.GET("someStruct", func(c *gin.Context) {
		var ss struct {
			Id     int
			Name   string
			Status string
		}
		ss.Id = 1
		ss.Name = "ganyu"
		ss.Status = "keqing"
		c.JSON(http.StatusOK, ss)
	})
	// 3 xml 响应
	r.GET("someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "XML",
		})
	})
	// yaml
	r.GET("someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "YAML",
		})
	})
	// protobuf
	r.GET("somePROTOBUF", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8000")
}
