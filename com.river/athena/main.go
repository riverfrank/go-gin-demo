package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/com.river/user"
)

func main() {
	//users := [2]User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})
	})

	//router.GET("/users", func(c *gin.Context) {
	//	c.JSON(200, users)
	//})

	//router.GET("/users", user.FindUsers)

	router.GET("/users/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		c.String(200, "The user id is  %s name is %s", id, name)
	})

	userRouter := router.Group("/users")
	userRouter.GET("", user.FindUsers)

	router.Run(":8080")

}
