package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
type PromotedIncome struct {
	income int
	user int
}
type date struct {
	date string
	data PromotedIncome
}
func main()  {
	fmt.Println("log")
	g := gin.Default()
	g.GET("/test", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"name":"zhuyuchen",
			"age":123,
		})
		return
	})
	g.Run(":9999")
}
