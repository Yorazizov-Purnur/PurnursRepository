package main

import "github.com/gin-gonic/gin"

type Users struct {
	Name    string
	Surname string
	Pasword string
	Login   string
}

var UserSlice = []Users{}

func main() {
	r := gin.Default()
	r.POST("/getuser", Register)
	r.GET("/list", List)
	r.Run(":1303")
}


func Register(c *gin.Context) {
	var shablon Users
	if shablon.Login == "" || shablon.Name == "" || shablon.Pasword == "" || shablon.Surname == "" {
		c.JSON(404, "empty field")
	} else {
		c.JSON(200, "succes")
		UserSlice = append(UserSlice, shablon)
	}
}

func List(c *gin.Context)  {
	c.JSON(200,UserSlice)
}