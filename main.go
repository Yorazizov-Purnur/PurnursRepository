package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/getuser")
	r.Run(":1303")
}