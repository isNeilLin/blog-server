package controllers

import (
	"github.com/gin-gonic/gin"
	"blog/models"
	"net/http"
	"strconv"
)

func Post(c *gin.Context) {
	title 		:= c.PostForm("title")
	content 	:= c.PostForm("content")
	summary 	:= c.PostForm("summary")
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
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"insert successed!",
			"post":		post,
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

func Update(c *gin.Context) {
	id			:= c.PostForm("id")
	title 		:= c.PostForm("title")
	content 	:= c.PostForm("content")
	summary 	:= c.PostForm("summary")
	str_publish := c.PostForm("publish")
	publish, _ 	:= strconv.ParseBool(str_publish)

	pid, err 	:= strconv.ParseUint(id, 10, 64)

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
	post.ID = uint(pid)
	err = post.Update()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
			"post":	post,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":		0,
			"message":	"update successed",
			"post":	post,
		})
	}
}

func Delete(c *gin.Context) {
	id 		:= c.PostForm("id")
	pid,err	:= strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":		1,
			"message":	err.Error(),
		})
		return
	}

	post 	:= &models.Post{}

	post.ID	= uint(pid)

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
