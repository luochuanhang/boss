package router

import (
	"boos/pkg/handler"

	"github.com/gin-gonic/gin"
)

//启动路由服务器
func Start() {
	r := gin.Default()
	//查
	r.GET("/resume", handler.Get)
	//增
	r.POST("/resume", handler.Post)
	//改
	r.PUT("/resume", handler.Put)
	//删除
	r.DELETE("/resume", handler.DELETE)
	//监听端口
	r.Run()
}
