package router

import (
	"blog/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 初始化
	r := gin.Default()

	// 获取文章列表
	r.GET("/get_all", controllers.GetAll)
	r.GET("/get_publish", controllers.GetPublish)
	r.GET("/get_private", controllers.GetPrivate)

	// 操作文章
	r.GET("/post", controllers.GetPostById)
	r.POST("/post", controllers.Post)
	r.POST("/update", controllers.Update)
	r.POST("/delete", controllers.Delete)

	// 标签相关
	r.GET("/tags",controllers.GetTags)
	r.POST("/create_tag",controllers.AddTag)
	r.POST("/update_tag",controllers.UpdateTag)
	r.POST("/delete_tag",controllers.DeleteTag)

	// 归档
	r.GET("/archive",controllers.GetArchive)

	r.Run()
}
