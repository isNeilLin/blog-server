package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func User(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}
func Welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "hello %s %s", firstname, lastname)
}
func FormPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	nick := c.DefaultPostForm("nick", "anonymous")
	message := c.PostForm("message")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"page":    page,
		"id":      id,
		"message": message,
		"nick":    nick,
	})
}
func Upload(c *gin.Context) {
	/* single File upload
	file, _ := c.FormFile("file")
	path := "./blog/static/" + file.Filename
	c.SaveUploadedFile(file,path)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	*/
	/* multiple files */
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		log.Println(file.Filename)
		path := "./blog/static/" + file.Filename
		c.SaveUploadedFile(file, path)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}