package controllers

import (
	"github.com/gin-gonic/gin"
	"blog/models"
	"net/http"
	"blog/utils"
	"strconv"
)

// 获取全部文章列表
func GetAll(c *gin.Context) {
	posts,err := models.GetAllPost()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"get list successed",
			"posts":	posts,
		})
	}
}

// 获取已发布文章
func GetPublish(c *gin.Context) {
	posts,err := models.GetPublishPost()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"get published list successed",
			"posts":	posts,
		})
	}
}

// 获取私密文章
func GetPrivate(c *gin.Context) {
	posts, err := models.GetPrivatePost()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"get published list successed",
			"posts":	posts,
		})
	}
}

// 创建文章
func Post(c *gin.Context) {
	title 		:= c.PostForm("title")
	content 	:= c.PostForm("content")
	summary 	:= c.PostForm("summary")
	tags		:= c.DefaultPostForm("tags","")
	str_publish := c.PostForm("publish")
	publish, _ 	:= strconv.ParseBool(str_publish)

	post := &models.Post{
		Title:		title,
		Content: 	content,
		Summary:	summary,
		Publish:	publish,
	}
	err := post.Insert()

	if err != nil {
		 c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}
	err = models.CreateTags(tags, post)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"insert successed!",
			"post":		post,
		})
	}
}

// 查看文章详情
func GetPostById(c *gin.Context) {
	id 		:= c.Query("id")
	pid,err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
	}
	post	:= &models.Post{}
	post.ID  = pid
	post,err =post.GetPostById()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code"		: 0,
			"message" 	: "query successed",
			"post"		: post,
		})
	}
}

// 更新文章
func Update(c *gin.Context) {
	id			:= c.PostForm("id")
	title 		:= c.PostForm("title")
	tags		:= c.PostForm("tags")
	content 	:= c.PostForm("content")
	summary 	:= c.PostForm("summary")
	str_publish := c.PostForm("publish")
	publish, _ 	:= strconv.ParseBool(str_publish)


	pid,err 	:= utils.StringToUint(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}
	post := &models.Post{
		Title:		title,
		Content:	content,
		Summary:	summary,
		Publish:	publish,
	}
	post.ID = pid
	err = post.Update()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}
	err = models.DeleteTagsFromPostId(pid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}
	err = models.CreateTags(tags,post)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":		0,
		"message":	"update successed",
		"post":	post,
	})
}

// 删除文章
func Delete(c *gin.Context) {
	id 		:= c.PostForm("id")
	pid,err	:= utils.StringToUint(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}

	post 	:= &models.Post{}
	post.ID	= pid
	err 	= post.Delete()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}
	err 	= models.DeleteTagsFromPostId(pid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":		0,
		"message":	"delete successed",
		"post":	post,
	})
}

// 获取标签列表 or 单个标签下的文章列表
func GetTags(c *gin.Context)  {
	// 标签ID
	id	:= c.DefaultQuery("id","")

	if id != "" {
		tid,err	:= utils.StringToUint(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":		1,
				"message":	err.Error(),
			})
			return
		}

		posts, err := models.GetPostsByTag(tid)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":		1,
				"message":	err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":		0,
				"message":	"get post successed",
				"posts":	posts,
			})
		}
	} else {
		tags, err := models.GetTags()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":		1,
				"message":	err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":		0,
				"message":	"get list successed",
				"tags":	tags,
			})
		}
	}
}

// 添加标签
func AddTag(c *gin.Context)  {
	name 	:= c.PostForm("name")
	color 	:= c.PostForm("color")
	tag 	:= &models.Tag{
		Name: 	name,
		Color:	color,
	}
	err 	:= tag.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"insert successed!",
			"tag":		tag,
		})
	}
}

// 更新标签
func UpdateTag(c *gin.Context)  {
	name 	:= c.PostForm("name")
	color 	:= c.PostForm("color")
	id 		:= c.PostForm("id")
	pid,err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}

	tag 	:= &models.Tag{
		Name: 	name,
		Color:	color,
	}
	tag.ID 	= pid
	err    	= tag.Update()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"update successed",
			"tag":	tag,
		})
	}
}

// 删除标签
func DeleteTag(c *gin.Context)  {
	id 		:= c.PostForm("id")
	pid,err	:= utils.StringToUint(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}

	tag 	:= &models.Tag{}
	tag.ID   = pid
	err 	 = tag.Delete()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"delete successed",
			"tag":	tag,
		})
	}
}

// 获取归档
func GetArchive(c *gin.Context)  {
	posts,err := models.GetPublishPost()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
	} else {
		result := map[int][]*models.Post{}
		for _,post := range posts {
			year := post.CreatedAt.Year()
			result[year] = append(result[year],post)
		}
		c.JSON(http.StatusOK,gin.H{
			"code":  	0,
			"message": 	"get post successed",
			"posts": 	 result,
		})
	}
}