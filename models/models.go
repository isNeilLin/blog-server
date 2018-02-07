package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"blog/utils"
	"blog/conf"
	"strings"
)

type Post struct {
	gorm.Model
	Title			string		`gorm:"not null"`
	Content 		string 		`gorm:"not null"`
	Summary 		string 		`gorm:"not null"`
	Publish			bool 		`gorm:"not null"`
	Status 			int 		`gorm:"default:1"`
}
type Tag struct {
	gorm.Model
	Name 			string 		`gorm:"not null"`
	Color 			string 		`gorm:"not null"`
}
type PostTag struct {
	gorm.Model
	PostId			uint		`gorm:"not null"`
	TagId			uint		`gorm:"not null"`
}

func (Post)TableName() string {
	return "post"
}

func (Tag)TableName() string {
	return "tag"
}

func (PostTag)TableName()string  {
	return "post_tag"
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	utils.InitConfig()
	db,err := gorm.Open(conf.LocalDB.Name,conf.LocalDB.Option)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Post{},&Tag{},&PostTag{})
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
		"summary":	post.Summary,
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

func (p *Post)GetPostById() (*Post, error){
	var post Post
	err := DB.First(&post, "id = ?", p.ID).Error
	return &post, err
}

func GetTags()([]map[string]interface{}, error)  {
	var (
		tags 	[]*Tag
		result 	[]map[string]interface{}
	)
	err := DB.Find(&tags).Error
	if err != nil {
		return result, err
	}
	for _,tag := range tags {
		posts, err := GetPostsByTag(tag.ID)
		if err != nil {
			return result, err
		} else {
			result = append(result,map[string]interface{}{
				"name": 	tag.Name,
				"color":	tag.Color,
				"posts":	posts,
			})
		}
	}
	return result, err
}

func GetPostsByTag(tagId uint) ([]*Post, error){
	var (
		posts 	[]*Post
	)
	err := DB.Joins("JOIN post_tag on post.id = post_tag.post_id").Where("post_tag.tag_id = ?",tagId).Find(&posts).Error
	return posts,err
}

func (tag *Tag)Insert() error {
	return DB.Create(tag).Error
}

func (tag *Tag)Update() error {
	return DB.Model(tag).Update(map[string]interface{}{
		"color": tag.Color,
		"name":  tag.Name,
	}).Error
}

func (tag *Tag)Delete() error {
	return DB.Delete(tag).Error
}

func (pt *PostTag)Insert() error  {
	return DB.Create(pt).Error
}

func CreateTags(tags string, post *Post) error {
	sliceTags := strings.Split(tags,",")
	for _,tag := range sliceTags {
		tagID,err := utils.StringToUint(tag)
		if err != nil {
			return err
		}
		pt := PostTag{
			PostId: post.ID,
			TagId: 	tagID,
		}
		err = pt.Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteTagsFromPostId(id uint) error {
	return DB.Delete(&PostTag{},"post_id = ?",id).Error
}