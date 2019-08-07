package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type user struct {
	Id   int
	Name string
}

func main() {
	r := gin.Default()

	r.GET("/hi", func(c *gin.Context) {
		fmt.Println("req.")
		u := user{1002, "user_name_1"}
		c.JSON(200, u)
	})

	r.POST("/register", func(c *gin.Context) {
		var u user
		err := c.BindJSON(&u)
		if err == nil {
			//no error
			fmt.Println(u.Id)
			fmt.Println(u.Name)

			c.JSON(200,gin.H{"message":"inserted"})
		} else {
			//error
			fmt.Println(err.Error())

			c.JSON(400, gin.H{"message":err.Error()})
		}
	})

	r.Run(":8070") // listen and serve on 0.0.0.0:8080
}
