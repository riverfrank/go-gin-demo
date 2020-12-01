package user

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func FindUsers(c *gin.Context) {

	users := [2]User{{ID: 1, Name: "river"}, {ID: 2, Name: "frank"}}

	c.JSON(200, users)
}
