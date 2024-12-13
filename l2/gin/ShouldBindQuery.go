package main

import "log"
import "github.com/gin-gonic/gin"

type Person2 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage1)
	route.Run(":8085")
}

func startPage1(c *gin.Context) {
	var person Person2
	if c.BindQuery(&person) == nil {
		log.Println("====== Only Bind Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}

//localhost:8085/testing?name=eason&address=xyz"
//$ curl -X POST "localhost:8085/testing?name=eason&address=xyz" --data 'name=ignore&address=ignore' -H "Content-Type:application/x-www-form-urlencoded"
