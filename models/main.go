package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"blog/utils"
	"blog/conf"
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