package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//配置路由
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{ //context.JSON返回JSON格式的数据 // H is a shortcut for map[string]interface{}  //type H map[string]any
			"message": "Hello world111!",
		})
	})
	r.GET("/getString", func(context *gin.Context) {
		context.String(http.StatusOK, "Try the string one, Hello world111!")
	})
	r.POST("/post", func(context *gin.Context) {
		context.String(200, "POST!!!!")
	})
	r.PUT("/put", func(context *gin.Context) {
		context.String(200, "PUT!!!!")
	})
	r.DELETE("/delete", func(context *gin.Context) {
		context.String(200, "DELETE!!!!!!!!!!!!!")
	})
	//启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()

	//fmt.Println("git test")
}
