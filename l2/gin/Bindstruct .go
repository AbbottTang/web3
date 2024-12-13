package main

import "github.com/gin-gonic/gin"

// 绑定表单数据至自定义结构体
type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)

	r.Run()
}

//注意：不支持以下格式结构体：总之, 目前仅支持没有 form 的嵌套结构体。
//type StructX struct {
//    X struct {} `form:"name_x"` // 有 form
//}
//
//type StructY struct {
//    Y StructX `form:"name_y"` // 有 form
//}
//
//type StructZ struct {
//    Z *StructZ `form:"name_z"` // 有 form
//}

//http://localhost:8080/getb?field_a=hello&field_b=world
//http://localhost:8080/getc?field_a=hello&field_c=world
//http://localhost:8080/getd?field_x=hello&field_d=world
