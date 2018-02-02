package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"blog/conf"
	"strconv"
)

type Post struct {
	gorm.Model
	Title			string		`gorm:"not null"`
	Content 		string 		`gorm:"not null"`
	Summary 		string 		`gorm:"not null"`
	Publish			bool 		`gorm:"not null"`
	Status 			int 		`gorm:"default:1"`
}

func (Post)TableName() string {
	return "post"
}

type Tag struct {
	gorm.Model
	Name 			string 		`gorm:"not null"`
	Color 			string 		`gorm:"not null"`
}

func (Tag)TableName() string {
	return "tag"
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	db,err := gorm.Open("mysql","root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Post{},&Tag{})
	DB = db
	return db
}

func (post *Post)Insert() error {
	return DB.Create(post).Error
}

func (post *Post)Update() error {
	return DB.Model(post).Updates(map[string]interface{}{
		"title":	post.Title,
		"content":	post.Content,
		"summary":	post.Status,
		"publish":	post.Publish,
	}).Error
}

func (post *Post)Delete() error {
	return DB.Delete(post).Error
}

func GetPublishPost() ([]*Post, error) {
	return getPostList(true)
}

func GetPrivatePost() ([]*Post, error)  {
	return getPostList(false)
}

func GetAllPost() ([]*Post, error) {
	var posts []*Post
	err := DB.Find(&posts).Error
	return posts, err
}

func getPostList(publish bool) ([]*Post, error) {
	var posts []*Post
	err := DB.Where("publish = ?", publish).Find(&posts).Error
	return posts, err
}

func GetPostById(id string) (*Post, error){
	pid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return nil, err
	}
	var post Post
	err = DB.First(&post, "id = ?", uint(pid)).Error
	return &post, err
}