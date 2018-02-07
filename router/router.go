package router

import (
	"blog/controllers"
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"time"
)

func InitRouter() {
	// 初始化
	r 		:= gin.Default()
	auth 	:= r.Group("auth")
	r.POST("/login", authMiddleWare.LoginHandler)
	// 获取文章列表
	r.GET("/get_publish", controllers.GetPublish)
	// 标签相关
	r.GET("/tags",controllers.GetTags)
	// 归档
	r.GET("/archive",controllers.GetArchive)

	auth.Use(authMiddleWare.MiddlewareFunc())
	{
		auth.GET("/get_all", controllers.GetAll)
		auth.GET("/get_private", controllers.GetPrivate)
		// 操作文章
		auth.GET("/post", controllers.GetPostById)
		auth.POST("/post", controllers.Post)
		auth.POST("/update", controllers.Update)
		auth.POST("/delete", controllers.Delete)
		auth.POST("/create_tag",controllers.AddTag)
		auth.POST("/update_tag",controllers.UpdateTag)
		auth.POST("/delete_tag",controllers.DeleteTag)
	}

	r.Run()
}

var authMiddleWare = &jwt.GinJWTMiddleware{
	Realm: 			"test zone",
	Key:			[]byte("the little prince"),
	Timeout:		time.Hour,
	MaxRefresh:		time.Hour,
	Authenticator:  func(user string, password string, c *gin.Context) (string, bool) {
		if user == "admin" && password == "admin" {
			return user, true
		}
		return user, false
	},
	Authorizator:   func(user string, c *gin.Context) bool {
		if user == "admin" {
			return true
		}
		return false
	},
	Unauthorized:	func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code": 	code,
			"message":	message,
		})
	},
	TokenLookup:	"header:Authorizator",
	TokenHeadName:	"Bearer",
	TimeFunc:		time.Now,
}
