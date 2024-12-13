package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type modelA struct {
	Foo string `form:"foo" json:"foo" binding:"required"`
	Bar int64  `form:"bar" json:"bar" binding:"required"`
}

type modelB struct {
	Hoge string `form:"hoge" json:"hoge" xml:"hoge" binding:"required"`
	Fuga int64  `form:"fuga" json:"fuga" xml:"fuga" binding:"required"`
}

func main() {
	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	//要想多次绑定，可以使用 c.ShouldBindBodyWith.

	r := gin.New()
	r.POST("/json", func(c *gin.Context) {
		fa, fb := modelA{}, modelB{}
		if aerr := c.ShouldBindBodyWith(&fa, binding.JSON); aerr == nil {
			c.JSON(http.StatusOK, gin.H{"form": fa})
		} else if berr := c.ShouldBindBodyWith(&fb, binding.JSON); berr == nil {
			c.JSON(http.StatusOK, gin.H{"form": fb})
		} else {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"errors": gin.H{
						"modelA": aerr.Error(),
						"modelB": berr.Error(),
					},
				},
			)
		}
	})
	r.Run(":8080")
}

// c.ShouldBind 使用了 c.Request.Body，不可重用。
//要想多次绑定，可以使用 c.ShouldBindBodyWith.

// // c.ShouldBind 使用了 c.Request.Body，不可重用。
//  if errA := c.ShouldBind(&objA); errA == nil {
//    c.String(http.StatusOK, `the body should be formA`)
//  // 因为现在 c.Request.Body 是 EOF，所以这里会报错。
//  } else if errB := c.ShouldBind(&objB); errB == nil {
//    c.String(http.StatusOK, `the body should be formB`)
//  } else {
//    ...
//  }

// // 读取 c.Request.Body 并将结果存入上下文。
//  if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
//    c.String(http.StatusOK, `the body should be formA`)
//  // 这时, 复用存储在上下文中的 body。
//  } else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
//    c.String(http.StatusOK, `the body should be formB JSON`)
//  // 可以接受其他格式
//  } else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
//    c.String(http.StatusOK, `the body should be formB XML`)
//  } else {
//    ...
//  }
