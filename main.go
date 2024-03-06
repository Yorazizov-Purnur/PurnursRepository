package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/getuser")
	r.GET("/list", List)
	r.Run(":1303")
}

func List(c *gin.Context)  {
	c.JSON(200,UserSlice)
}