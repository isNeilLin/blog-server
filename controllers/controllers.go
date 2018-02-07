package controllers

import (
	"github.com/gin-gonic/gin"
	"blog/models"
	"net/http"
	"blog/utils"
	"strconv"
)

func Post(c *gin.Context) {
	title 		:= c.PostForm("title")
	content 	:= c.PostForm("content")
	summary 	:= c.PostForm("summary")
	tags		:= c.PostForm("tags")
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
			"post":		post,
		})
	} else {
		models.CreateTags(tags, post)
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"insert successed!",
			"post":		post,
		})
	}
}

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
	} else {
		err = models.DeleteTagsFromPostId(pid)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":		1,
				"message":	err.Error(),
			})
		}
		err = models.CreateTags(tags,post)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":		1,
				"message":	err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"update successed",
			"post":	post,
		})
	}
}

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
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"delete successed",
			"post":	post,
		})
	}

}

func GetTags(c *gin.Context)  {
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

func GetPostByTag(c *gin.Context)  {
	id		:= c.Query("tag_id")
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
}