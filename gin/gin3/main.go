package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)

	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
	}

	router.Run(":8080")
}

func loginEndpoint(c *gin.Context) {

}
func submitEndpoint(c *gin.Context) {

}
