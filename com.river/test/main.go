package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint64
	Name string
}

func main() {
	users := [2]User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	r.GET("/users/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		c.String(200, "The user id is  %s name is %s", id, name)
	})
	r.Run(":8080")
}
